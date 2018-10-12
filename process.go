package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jpnewman/urlinfo/profiling"

	"github.com/motemen/go-loghttp"
	"github.com/sirupsen/logrus"
)

type processURLsArgs struct {
	httpTimeoutMilliseconds time.Duration
	numberOfWorkers         int
	getHeadOny              bool
	dontFollowRedirects     bool
	dryRun                  bool
}

type httpRequestArgs struct {
	url     string
	options *processURLsArgs
}

type httpResponse struct {
	url           string
	statusCode    int
	contentLength int64
	headers       map[string][]string
	body          string
	errs          []error
}

func createHTTPClient(args *httpRequestArgs) *http.Client {
	client := &http.Client{
		Transport: &loghttp.Transport{
			LogRequest: func(req *http.Request) {
				Logger.WithFields(logrus.Fields{
					"method": req.Method,
					"url":    req.URL,
				}).Debug("HTTP Request")
			},
			LogResponse: func(resp *http.Response) {
				Logger.WithFields(logrus.Fields{
					"code": resp.StatusCode,
					"url":  resp.Request.URL,
				}).Debug("HTTP Response")
			},
		},
		Timeout: args.options.httpTimeoutMilliseconds,
	}

	if args.options.dontFollowRedirects {
		client.CheckRedirect =
			func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
	}

	return client
}

func httpRequest(args *httpRequestArgs, client *http.Client) (*http.Response, error) {
	defer profiling.TimeElapsed("HTTP Request Time")(Report.PrintMessage)
	defer profiling.TimeElapsed(fmt.Sprintf("HTTP Request Time: %s", args.url))(LogPrintInfo)

	var resp *http.Response
	var err error
	if args.options.getHeadOny {
		resp, err = client.Head(args.url)
	} else {
		resp, err = client.Get(args.url)
	}

	return resp, err
}

func getHTTPResponseBody(resp *http.Response) (string, error) {
	if resp == nil {
		return "", nil
	}

	if resp.Body == nil {
		return "", errors.New("HTTP Response Body not defined")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if b == nil {
		return "", errors.New("HTTP Response read Body error")
	}

	return string(b), err
}

func getHTTPResponse(args *httpRequestArgs) *httpResponse {
	httpResp := new(httpResponse)
	client := createHTTPClient(args)

	httpResp.url = args.url

	if args.options.dryRun {
		time.Sleep(args.options.httpTimeoutMilliseconds)
		httpResp.errs = append(httpResp.errs, errors.New("Dry-Run Mode"))
		return httpResp
	}

	resp, err := httpRequest(args, client)

	if err == nil {
		defer resp.Body.Close()
	} else {
		Logger.Error(err)
		httpResp.errs = append(httpResp.errs, err)
	}

	if resp == nil {
		return httpResp
	}

	httpResp.statusCode = resp.StatusCode
	httpResp.contentLength = resp.ContentLength
	httpResp.headers = resp.Header

	body, bodyErr := getHTTPResponseBody(resp)
	httpResp.body = body
	httpResp.errs = append(httpResp.errs, bodyErr)

	return httpResp
}

func worker(jobs <-chan *httpRequestArgs, results chan<- *httpResponse) {
	for j := range jobs {
		var errs []error
		httpResp := getHTTPResponse(j)
		errs = append(errs, httpResp.errs...)

		results <- httpResp
	}
}

func processURLs(urls map[string][]lineDetail, args *processURLsArgs) {
	urlsCount := len(urls)

	Report.PrintHeaderf("Processing URLs %d", urlsCount)
	Report.PrintSubHeaderf("Workers: %d", args.numberOfWorkers)

	jobs := make(chan *httpRequestArgs, urlsCount)
	results := make(chan *httpResponse, urlsCount)
	defer close(results)

	for w := 1; w <= args.numberOfWorkers; w++ {
		go worker(jobs, results)
	}

	for key := range urls {
		jobs <- &httpRequestArgs{
			url:     key,
			options: args,
		}
	}
	close(jobs)

	errorCount := 0
	for j := 0; j < urlsCount; j++ {
		ret := <-results
		errorCount += printOutput(args, ret)
	}

	Report.PrintHeaderf("Processed %d URLs with %d Errors", urlsCount, errorCount)
}
