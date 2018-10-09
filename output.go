package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"

	humanize "github.com/dustin/go-humanize"
	logging "github.com/jpnewman/urlinfo/logging"
	"github.com/sirupsen/logrus"
)

// PrintURLInfo Print URL Info
func PrintURLInfo(resp *httpResponse) {
	if resp.statusCode == http.StatusOK {
		Report.PrintOKf("%s : %d", resp.url, resp.statusCode)
		PrintHTTPHeaders(resp)
	}
}

// PrintStats Print Stats
func PrintStats() {
	// TODO: Implement
}

// PrintHTTPHeaders Print HTTP Headers
func PrintHTTPHeaders(resp *httpResponse) {
	var buffer bytes.Buffer

	var keys []string
	for k := range resp.headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		buffer.WriteString(fmt.Sprintf("%v: %v\n", k, resp.headers[k]))
	}

	Report.PrintCode(buffer.String())
}

func printOutput(args processURLsArgs, ret *httpResponse) int {
	Report.PrintSubHeader(ret.url)

	errorCount := 0
	for _, e := range ret.errs {
		if e != nil {
			Report.PrintError(e)
			logging.Logger.Error(e)

			errorCount++
		}
	}

	if logging.RootLogger.GetLevel() == logrus.DebugLevel {
		PrintURLInfo(ret)
	}

	if args.getHeadOny {
		Report.PrintMessagef("Header ContentLength (Octets): %d\n", ret.contentLength)
	} else if errorCount == 0 {
		bodyLen := uint64(len(ret.body))
		Report.PrintMessagef("Body Length: %d (%s)", bodyLen, humanize.Bytes(bodyLen))
	}

	PrintStats()

	return errorCount
}
