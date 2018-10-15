package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

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

type httpResponse struct {
	url           string
	statusCode    int
	contentLength int64
	headers       map[string][]string
	requestTime   time.Duration
	body          string
	errs          []error
}

func createHTTPClient(httpTimeoutMilliseconds time.Duration, dontFollowRedirects bool) *http.Client {
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
		Timeout: httpTimeoutMilliseconds,
	}

	if dontFollowRedirects {
		client.CheckRedirect =
			func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
	}

	return client
}

func httpRequest(url string, args *processURLsArgs, client *http.Client) (*http.Response, time.Duration, error) {
	var resp *http.Response
	var err error

	startTime := time.Now()
	if args.getHeadOny {
		resp, err = client.Head(url)
	} else {
		resp, err = client.Get(url)
	}
	requestTime := time.Since(startTime)

	return resp, requestTime, err
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

func getHTTPResponse(url string, args *processURLsArgs, client *http.Client) *httpResponse {
	httpResp := new(httpResponse)
	httpResp.url = url

	if args.dryRun {
		time.Sleep(args.httpTimeoutMilliseconds)
		httpResp.errs = append(httpResp.errs, errors.New("Dry-Run Mode"))
		return httpResp
	}

	resp, requestTime, err := httpRequest(url, args, client)

	if err != nil {
		Logger.Error(err)
		httpResp.errs = append(httpResp.errs, err)
	}

	if resp == nil {
		return httpResp
	}

	httpResp.statusCode = resp.StatusCode
	httpResp.contentLength = resp.ContentLength
	httpResp.headers = resp.Header
	httpResp.requestTime = requestTime

	body, bodyErr := getHTTPResponseBody(resp)
	httpResp.body = body
	httpResp.errs = append(httpResp.errs, bodyErr)

	defer resp.Body.Close()

	return httpResp
}

func worker(jobs <-chan string, results chan<- *httpResponse, args *processURLsArgs, client *http.Client) {
	for j := range jobs {
		var errs []error
		httpResp := getHTTPResponse(j, args, client)
		errs = append(errs, httpResp.errs...)

		results <- httpResp
	}
}

func processURLs(urls map[string][]lineDetail, args *processURLsArgs, client *http.Client) {
	urlsCount := len(urls)

	Report.PrintHeaderf("Processing URLs %d", urlsCount)
	Report.PrintSubHeaderf("Workers: %d", args.numberOfWorkers)

	jobs := make(chan string, urlsCount)
	results := make(chan *httpResponse, urlsCount)
	defer close(results)

	for w := 1; w <= args.numberOfWorkers; w++ {
		go worker(jobs, results, args, client)
	}

	for key := range urls {
		jobs <- key
	}
	close(jobs)

	errorCount := 0
	for j := 0; j < urlsCount; j++ {
		ret := <-results
		errorCount += printOutput(args, ret)
	}

	Report.PrintHeaderf("Processed %d URLs with %d Errors", urlsCount, errorCount)
}
