package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jpnewman/urlinfo/profiling"

	"github.com/motemen/go-loghttp"
	"github.com/sirupsen/logrus"

	logging "github.com/jpnewman/urlinfo/logging"
)

type processURLsArgs struct {
	httpTimeoutMilliseconds int
	numberOfWorkers         int
	getHeadOny              bool
	dontFollowRedirects     bool
	dryRun                  bool
}

type getHTTPArgs struct {
	url     string
	options processURLsArgs
}

type httpResponse struct {
	url           string
	statusCode    int
	contentLength int64
	headers       map[string][]string
	body          string
	errs          []error
}

func createHTTPClient(args getHTTPArgs, timeout time.Duration) *http.Client {
	client := &http.Client{
		Transport: &loghttp.Transport{
			LogRequest: func(req *http.Request) {
				logging.Logger.WithFields(logrus.Fields{
					"method": req.Method,
					"url":    req.URL,
				}).Debug("HTTP Request")
			},
			LogResponse: func(resp *http.Response) {
				logging.Logger.WithFields(logrus.Fields{
					"code": resp.StatusCode,
					"url":  resp.Request.URL,
				}).Debug("HTTP Response")
			},
		},
		Timeout: timeout,
	}

	if args.options.dontFollowRedirects {
		client.CheckRedirect =
			func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
	}

	return client
}

func httpRequest(args getHTTPArgs, client *http.Client) (*http.Response, error) {
	defer profiling.Elapsed("HTTP Request Time")()

	var resp *http.Response
	var err error
	if args.options.getHeadOny {
		resp, err = client.Head(args.url)
	} else {
		resp, err = client.Get(args.url)
	}

	if err != nil {
		logging.Logger.Error(err)
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

func getHTTP(args getHTTPArgs) *httpResponse {
	timeout := time.Duration(time.Duration(args.options.httpTimeoutMilliseconds) * time.Millisecond)

	httpResp := new(httpResponse)
	client := createHTTPClient(args, timeout)

	httpResp.url = args.url

	if args.options.dryRun {
		time.Sleep(timeout)
		httpResp.errs = append(httpResp.errs, errors.New("Dry-Run Mode"))
		return httpResp
	}

	resp, err := httpRequest(args, client)
	httpResp.errs = append(httpResp.errs, err)

	if err == nil {
		defer resp.Body.Close()
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

func worker(jobs <-chan getHTTPArgs, results chan<- httpResponse) {
	for j := range jobs {
		var errs []error
		httpResp := getHTTP(j)
		errs = append(errs, httpResp.errs...)

		results <- *httpResp
	}
}

func processURLs(urls map[string][]lineDetails, args processURLsArgs) {
	urlsCount := len(urls)

	Report.PrintHeaderf("Processing URLs %d", urlsCount)
	Report.PrintSubHeaderf("Workers: %d", args.numberOfWorkers)

	jobs := make(chan getHTTPArgs, urlsCount)
	results := make(chan httpResponse, urlsCount)
	defer close(results)

	for w := 1; w <= args.numberOfWorkers; w++ {
		go worker(jobs, results)
	}

	for key := range urls {
		jobs <- getHTTPArgs{
			url:     key,
			options: args,
		}
	}
	close(jobs)

	errorCount := 0
	for j := 0; j < urlsCount; j++ {
		ret := <-results
		errorCount += printOutput(args, &ret)
	}

	Report.PrintHeaderf("Processed %d URLs with %d Errors", urlsCount, errorCount)
}
