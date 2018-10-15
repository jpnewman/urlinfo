package main

import (
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testDefaultURL = "http://localhost:123456"
var testDefaultTimeout = 3000

func TestCreateHTTPClient(t *testing.T) {
	args := helperCreateProcessURLsArgs(testDefaultTimeout)
	client := createHTTPClient(args.httpTimeoutMilliseconds, args.dontFollowRedirects)

	var httpClient *http.Client
	assert.IsType(t, httpClient, client)

	timeoutDuration := time.Duration(time.Duration(testDefaultTimeout) * time.Millisecond)
	assert.Equal(t, timeoutDuration, client.Timeout)
}

func TestHTTPRequest_Localhost(t *testing.T) {
	args := helperCreateProcessURLsArgs(testDefaultTimeout)
	client := createHTTPClient(args.httpTimeoutMilliseconds, args.dontFollowRedirects)
	resp, _, err := httpRequest(testDefaultURL, args, client)

	assert.NotNil(t, err)
	assert.Nil(t, resp)

	var urlError *url.Error
	assert.IsType(t, urlError, err)

	var httpResponse *http.Response
	assert.IsType(t, httpResponse, resp)
}

func TestHTTPRequest_LocalhostMock(t *testing.T) {
	args := helperCreateProcessURLsArgs(testDefaultTimeout)
	client := mockCreateHTTPClient(t, args.httpTimeoutMilliseconds, 401, 0)
	resp, _, err := httpRequest(testDefaultURL, args, client)

	assert.Nil(t, err)
	assert.NotNil(t, resp)

	var httpResponse *http.Response
	assert.IsType(t, httpResponse, resp)

	assert.Equal(t, resp.StatusCode, 401)
}

// FIXME: Does not fail with timeout as expected. Due to a possible issue with the implemention of mockCreateHTTPClient.
// Fix or remove and replace with TestHTTPRequest_ClientTimeout.
func TestHTTPRequest_MockTimeout(t *testing.T) {
	t.Skip("Skipping test as it's not working as expected.")

	args := helperCreateProcessURLsArgs(testDefaultTimeout)
	responseTimeMillisecond := time.Duration(time.Duration(testDefaultTimeout*2) * time.Millisecond)

	assert.True(t, args.httpTimeoutMilliseconds < responseTimeMillisecond)

	client := mockCreateHTTPClient(t, args.httpTimeoutMilliseconds, 200, responseTimeMillisecond)
	resp, _, err := httpRequest(testDefaultURL, args, client)

	assert.Nil(t, err)
	assert.NotNil(t, resp)

	var httpResponse *http.Response
	assert.IsType(t, httpResponse, resp)

	netErr, ok := err.(net.Error)
	assert.True(t, ok && netErr.Timeout())
}

func TestHTTPRequest_ClientTimeout(t *testing.T) {
	responseTimeMillisecond := time.Duration(time.Duration(testDefaultTimeout*2) * time.Millisecond)

	server := newHTTPTestServer(t, responseTimeMillisecond)
	url := server.URL

	args := helperCreateProcessURLsArgs(testDefaultTimeout)

	client := createHTTPClient(args.httpTimeoutMilliseconds, args.dontFollowRedirects)
	resp, _, err := httpRequest(url, args, client)

	assert.NotNil(t, err)
	assert.Nil(t, resp)

	netErr, ok := err.(net.Error)
	assert.True(t, ok && netErr.Timeout())
}

func BenchmarkProcessURLs(b *testing.B) {
	b.ReportAllocs()

	urls := helperCreateLineDetails()
	args := helperCreateProcessURLsArgs(200)
	args.dryRun = true
	args.numberOfWorkers = 5

	Report.SetFormatter("none")

	client := createHTTPClient(args.httpTimeoutMilliseconds, args.dontFollowRedirects)

	for n := 0; n < b.N; n++ {
		processURLs(urls, args, client)
	}
}
