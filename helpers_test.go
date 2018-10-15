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

func helperCreateTestHTTPRequestArgs(url string) *getHTTPArgs {
	return &getHTTPArgs{
		url:     url,
		options: helperCreateProcessURLsArgs(),
	}
}

func helperCreateLineDetails() map[string][]lineDetails {
	urls := make(map[string][]lineDetails)

	for i, u := range testDefaultURLs {
		urls[u] = append(urls[u], lineDetails{
			lineNumber: i,
			line:       u,
			comment:    "",
		})
	}

	return urls
}
