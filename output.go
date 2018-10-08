package main

import (
	humanize "github.com/dustin/go-humanize"
	logging "github.com/jpnewman/urlinfo/logging"
	"github.com/sirupsen/logrus"
)

func printOutput(args processURLsArgs, ret getHTTPResult) int {
	Report.PrintSubHeader(ret.url)

	errorCount := 0
	for _, e := range ret.errs {
		if e != nil {
			Report.PrintError(e)
			logging.Logger.Error(e)

			errorCount++
		}
	}

	if args.getHeadOny {
		if ret.resp != nil {
			Report.PrintMessagef("Header ContentLength (Octets): %d\n", ret.resp.ContentLength)
		}
	} else if errorCount == 0 {
		bodyLen := uint64(len(ret.body))
		Report.PrintMessagef("Body Length: %d (%s)", bodyLen, humanize.Bytes(bodyLen))
	}

	if logging.RootLogger.GetLevel() == logrus.DebugLevel {
		Report.PrintURLInfo(ret.resp)
	}

	return errorCount
}
