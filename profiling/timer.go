package profiling

import (
	"time"

	logging "github.com/jpnewman/urlinfo/logging"
)

// Elapsed Elapsed Time, call with defer
func Elapsed(msg string) func() {
	startTime := time.Now()
	return func() {
		logging.Logger.Infof("%s: %s", msg, time.Since(startTime))
	}
}
