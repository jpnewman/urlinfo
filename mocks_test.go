package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Based On: http://hassansin.github.io/Unit-Testing-http-client-in-Go
// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestHTTPClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestHTTPClient(fn RoundTripFunc, timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
		Timeout:   timeout,
	}
}

func mockCreateHTTPClient(t *testing.T, timeout time.Duration, statusCode int, responseTimeMillisecond time.Duration) *http.Client {
	return NewTestHTTPClient(
		func(req *http.Request) *http.Response {
			resp := &http.Response{
				StatusCode: statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`TEST`)),
				Header:     make(http.Header),
			}

			time.Sleep(responseTimeMillisecond)

			return resp
		},
		timeout)
}

func newHTTPTestServer(t *testing.T, responseTimeMillisecond time.Duration) *httptest.Server {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseTimeMillisecond * 2)
		w.WriteHeader(http.StatusOK)
	})

	return httptest.NewServer(http.TimeoutHandler(handlerFunc, responseTimeMillisecond, "Server Timeout"))
}
