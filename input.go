package main

import (
	"bufio"
	"net/url"
	"os"
	"regexp"
	"strings"

	logging "github.com/jpnewman/urlinfo/logging"
)

type lineError struct {
	err        error
	lineNumber int
}

type lineDetails struct {
	lineNumber int
	line       string
	comment    string
}

func parseURLFileLine(line string, lineNum int, urls map[string][]lineDetails) []lineError {
	var errs []lineError

	line = strings.TrimSpace(line)
	lineComment := ""
	urlCommentLine := strings.SplitN(line, "#", 2)

	if len(urlCommentLine) == 2 {
		lineComment = urlCommentLine[1]
	}

	lineDetails := lineDetails{
		lineNumber: lineNum,
		line:       line,
		comment:    lineComment,
	}

	lineURL := strings.TrimSpace(urlCommentLine[0])
	u, err := url.ParseRequestURI(lineURL)
	if err != nil {
		errs = append(errs, lineError{
			err:        err,
			lineNumber: lineNum,
		})
	} else {
		key := strings.ToLower(u.String())
		_, contains := urls[key]
		if contains {
			urls[key] = append(urls[key], lineDetails)
		} else {
			urls[key] = append(urls[key], lineDetails)
		}
	}

	return errs
}

func readURLFile(path *string, count int) (map[string][]lineDetails, []lineError) {
	Report.PrintSubHeaderf("Parsing URL file: %s", *path)

	var lineRegEx = regexp.MustCompile(`(^\s*#.*$|^\s*$)`)
	urls := make(map[string][]lineDetails)
	var errs []lineError

	file, err := os.Open(*path)
	if err != nil {
		logging.Logger.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		var line = scanner.Text()
		if !lineRegEx.MatchString(line) {
			lineErr := parseURLFileLine(line, lineNum, urls)
			errs = append(errs, lineErr...)
		}
	}

	if err := scanner.Err(); err != nil {
		logging.Logger.Fatal(err)
	}

	return urls, errs
}

func printURLList(urls map[string][]lineDetails) {
	var printableList []string

	for key := range urls {
		printableList = append(printableList, key)
	}

	Report.PrintList(printableList)
}

func printURLs(urls map[string][]lineDetails) {
	Report.PrintMessagef("Found URLs (%d): -\n", len(urls))
	printURLList(urls)
}

func printURLDuplicates(urls map[string][]lineDetails) {
	dups := make(map[string][]lineDetails)

	for key, val := range urls {
		if len(val) > 1 {
			dups[key] = val
		}
	}

	Report.PrintMessagef("Duplicate URLs (%d): -\n", len(dups))

	for key, val := range dups {
		Report.PrintMessagef("- %s\n", key)
		for _, v := range val {
			Report.PrintMessagef("  [%d] %s\n", v.lineNumber, v.line)
		}
	}
}

func printURLErrors(errs []lineError) {
	if len(errs) > 0 {
		Report.PrintMessagef("Error URLs (%d): -\n", len(errs))
		for _, e := range errs {
			Report.PrintMessagef("- [%d] %s\n", e.lineNumber, e.err)
		}
	}
}

func printFileDetails(urls map[string][]lineDetails, errs []lineError) {
	printURLs(urls)
	printURLDuplicates(urls)
	printURLErrors(errs)
}
