package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/motemen/go-loghttp"
	"github.com/sirupsen/logrus"

	logging "github.com/jpnewman/urlinfo/logging"
)

type processURLsArgs struct {
	httpTimeoutSeconds  int
	numberOfWorkers     int
	getHeadOny          bool
	dontFollowRedirects bool
	dryRun              bool
}

type getHTTPArgs struct {
	url     string
	options processURLsArgs
}

type getHTTPResult struct {
	url  string
	resp *http.Response
	body string
	errs []error
}

// getHTTP - The caller should close http.Response.Body
func getHTTP(args getHTTPArgs) (*http.Response, error) {
	timeout := time.Duration(time.Duration(args.options.httpTimeoutSeconds) * time.Second)
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

	if args.options.dryRun {
		time.Sleep(timeout)
		return nil, errors.New("Dry-Run Mode")
	}

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

// getHTTPResponseBody - The caller should close http.Response.Body
func getHTTPResponseBody(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}

	if resp.Body == nil {
		return nil, errors.New("HTTP Response Body not defined")
	}

	b, err := ioutil.ReadAll(resp.Body)

	return b, err
}

func worker(jobs <-chan getHTTPArgs, results chan<- getHTTPResult) {
	for j := range jobs {
		var errs []error
		resp, err := getHTTP(j)
		errs = append(errs, err)

		var body []byte
		var bodyErr error
		if !j.options.getHeadOny {
			body, bodyErr = getHTTPResponseBody(resp)
			errs = append(errs, bodyErr)
		}

		ret := new(getHTTPResult)
		ret.url = j.url
		ret.resp = resp
		ret.body = string(body)
		ret.errs = errs
		results <- *ret

		// Closing http.Response.Body
		defer func() {
			if resp != nil {
				defer resp.Body.Close()
			}
		}()
	}
}

func processURLs(urls map[string][]lineDetails, args processURLsArgs) {
	urlsCount := len(urls)

	Report.PrintHeaderf("Processing URLs %d", urlsCount)
	Report.PrintSubHeaderf("Workers: %d", args.numberOfWorkers)

	jobs := make(chan getHTTPArgs, urlsCount)
	results := make(chan getHTTPResult, urlsCount)
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
		errorCount += printOutput(args, ret)
	}

	Report.PrintHeaderf("Processed %d URLs with %d Errors", urlsCount, errorCount)
}
