package profiling

import (
	"io/ioutil"
	"testing"
)

func testTimeElapseWriter(s string) {
	out := ioutil.Discard
	out.Write([]byte(s))
}

func TestTimeElapsed_NilWriters(t *testing.T) {
	TimeElapsed("Test")(nil, nil)

	// NOTE: Not Testable
}

func TestTimeElapsed_TimeElapseWriters(t *testing.T) {
	TimeElapsed("Test")(testTimeElapseWriter, testTimeElapseWriter)

	// NOTE: Not Testable
}

func BenchmarkTimeElapsed(b *testing.B) {
	TimeElapsed("Test")(testTimeElapseWriter, testTimeElapseWriter)
}
