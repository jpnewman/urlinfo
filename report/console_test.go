package report

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleSetOutput_stdout(t *testing.T) {
	var bw *bufio.Writer
	assert.IsType(t, consoleFormatter.Out, bw)

	consoleFormatter.SetOutput(os.Stdout)
	assert.IsType(t, consoleFormatter.Out, os.Stdout)

	var bb bytes.Buffer
	writer := bufio.NewWriter(&bb)
	consoleFormatter.SetOutput(writer)
}

func TestConsolePrintMessage_EmptyString(t *testing.T) {
	s := ""
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	consoleFormatter.PrintMessage(s)
	assert.Equal(t, "\n", bb.String())
}

func BenchmarkConsolePrintMessage_EmptyString(b *testing.B) {
	b.ReportAllocs()

	s := ""
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	for n := 0; n < b.N; n++ {
		consoleFormatter.PrintMessage(s)
	}
}

func TestConsolePrintMessage_TextWithoutNewline(t *testing.T) {
	s := "TEST"
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	consoleFormatter.PrintMessage(s)
	assert.Equal(t, fmt.Sprintf("%s\n", s), bb.String())
}

func TestConsolePrintMessage_TextWithNewline(t *testing.T) {
	s := "TEST\n"
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	consoleFormatter.PrintMessage(s)
	assert.Equal(t, s, bb.String())
}

// FIXME: Check or remove ANSI escape code and under-/over-line.
// func TestConsolePrintHeader_TextWithoutNewline(t *testing.T) {
// 	s := "TEST"
// 	var bb bytes.Buffer
// 	consoleFormatter.SetOutput(&bb)

// 	consoleFormatter.PrintHeader(s)

// 	assert.Equal(t, fmt.Sprintf("%s\n", s), bb.String())
// }

func BenchmarkConsolePrintHeader_EmptyString(b *testing.B) {
	b.ReportAllocs()

	s := ""
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	for n := 0; n < b.N; n++ {
		consoleFormatter.PrintHeader(s)
	}
}
