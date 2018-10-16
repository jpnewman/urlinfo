package main

// Test Helpers

import (
	"time"
)

var testDefaultURLs = [...]string{
	"https://google.com",
	"https://microsoft.com",
	"https://apple.com",
	"https://amazon.co.uk",
	"https://github.com",
}

func helperCreateProcessURLsArgs(httpTimeout int) *processURLsArgs {
	timeout := time.Duration(time.Duration(httpTimeout) * time.Millisecond)

	return &processURLsArgs{
		httpTimeoutMilliseconds: timeout,
		numberOfWorkers:         5,
		getHeadOny:              false,
		dontFollowRedirects:     false,
		dryRun:                  false,
	}
}

func helperCreateLineDetails() mapURLs {
	urls := make(mapURLs)

	for i, u := range testDefaultURLs {
		urls[u] = append(urls[u], lineDetail{
			lineNumber: i,
			line:       u,
			comment:    "",
		})
	}

	return urls
}
