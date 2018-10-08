package report

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var consoleFormatter *ConsoleFormatter

func setup() {
	fmt.Println("Test Setup...")

	var bb bytes.Buffer
	writer := bufio.NewWriter(&bb)

	consoleFormatter = &ConsoleFormatter{
		Out: writer,
	}
}

func teardown() {
	fmt.Println("Test Teardown...")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestSetOutput_stdout(t *testing.T) {
	var bw *bufio.Writer
	assert.IsType(t, consoleFormatter.Out, bw)

	consoleFormatter.SetOutput(os.Stdout)
	assert.IsType(t, consoleFormatter.Out, os.Stdout)

	var bb bytes.Buffer
	writer := bufio.NewWriter(&bb)
	consoleFormatter.SetOutput(writer)
}

func TestPrintMessage_EmptyString(t *testing.T) {
	s := ""
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	consoleFormatter.PrintMessage(s)
	assert.Equal(t, "\n", bb.String())
}

func BenchmarkPrintMessage_EmptyString(b *testing.B) {
	b.ReportAllocs()

	s := ""
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	for n := 0; n < b.N; n++ {
		consoleFormatter.PrintMessage(s)
	}
}

func TestPrintMessage_TextWithoutNewline(t *testing.T) {
	s := "TEST"
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	consoleFormatter.PrintMessage(s)
	assert.Equal(t, fmt.Sprintf("%s\n", s), bb.String())
}

func TestPrintMessage_TextWithNewline(t *testing.T) {
	s := "TEST\n"
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	consoleFormatter.PrintMessage(s)
	assert.Equal(t, s, bb.String())
}

// FIXME: Check or remove ANSI escape code and under-/over-line.
// func TestPrintHeader_TextWithoutNewline(t *testing.T) {
// 	s := "TEST"
// 	var bb bytes.Buffer
// 	consoleFormatter.SetOutput(&bb)

// 	consoleFormatter.PrintHeader(s)

// 	assert.Equal(t, fmt.Sprintf("%s\n", s), bb.String())
// }

func BenchmarkPrintHeader_EmptyString(b *testing.B) {
	b.ReportAllocs()

	s := ""
	var bb bytes.Buffer
	consoleFormatter.SetOutput(&bb)

	for n := 0; n < b.N; n++ {
		consoleFormatter.PrintHeader(s)
	}
}
