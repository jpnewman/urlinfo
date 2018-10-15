package main

// Test Helpers

var testDefaultURLs = [...]string{
	"https://google.com",
	"https://microsoft.com",
	"https://apple.com",
	"https://amazon.co.uk",
	"https://github.com",
}

func helperCreateProcessURLsArgs() processURLsArgs {
	return processURLsArgs{
		numberOfWorkers:     5,
		getHeadOny:          false,
		dontFollowRedirects: false,
		dryRun:              false,
	}
}

func helperCreateTestHTTPRequestArgs(url string) *httpRequestArgs {
	return &httpRequestArgs{
		url:     url,
		options: helperCreateProcessURLsArgs(),
	}
}

func helperCreateLineDetails() map[string][]lineDetail {
	urls := make(map[string][]lineDetail)

	for i, u := range testDefaultURLs {
		urls[u] = append(urls[u], lineDetail{
			lineNumber: i,
			line:       u,
			comment:    "",
		})
	}

	return urls
}
