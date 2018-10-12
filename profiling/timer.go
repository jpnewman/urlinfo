package profiling

import (
	"fmt"
	"time"
)

// TimeElapseWriter Time Elapse Writer
type TimeElapseWriter func(string)

// TimeElapsed TimeElapsed Time, call with a defer
func TimeElapsed(msg string) func(a ...TimeElapseWriter) {
	startTime := time.Now()
	return func(a ...TimeElapseWriter) {
		s := fmt.Sprintf("%s: %s\n", msg, time.Since(startTime))
		for _, fn := range a {
			if fn == nil {
				fmt.Print(s)
				continue
			}
			fn(s)
		}
	}
}
