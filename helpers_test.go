package main

// Test Helpers

import (
	"time"
)

func helperCreateTestHTTPRequestArgs(url string, httpTimeout int) *httpRequestArgs {
	timeout := time.Duration(time.Duration(httpTimeout) * time.Millisecond)

	return &httpRequestArgs{
		url: url,
		options: &processURLsArgs{
			httpTimeoutMilliseconds: timeout,
			numberOfWorkers:         5,
			getHeadOny:              false,
			dontFollowRedirects:     false,
			dryRun:                  false,
		},
	}
}
