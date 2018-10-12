package profiling

import (
	"fmt"
	"io"
	"time"
)

// Elapsed Elapsed Time, call with a defer
// TODO: Get formatted io.writer for logrus or change this function.
func Elapsed(msg string) func([]io.Writer) {
	startTime := time.Now()
	return func(wa []io.Writer) {
		s := fmt.Sprintf("%s: %s\n", msg, time.Since(startTime))
		for _, w := range wa {
			if w != nil {
				io.WriteString(w, s)
			} else {
				fmt.Print(s)
			}
		}
	}
}
