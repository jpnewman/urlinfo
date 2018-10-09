package report

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdownSetOutput_stdout(t *testing.T) {
	var bw *bufio.Writer
	assert.IsType(t, markdownFormatter.Out, bw)

	markdownFormatter.SetOutput(os.Stdout)
	assert.IsType(t, markdownFormatter.Out, os.Stdout)

	var bb bytes.Buffer
	writer := bufio.NewWriter(&bb)
	markdownFormatter.SetOutput(writer)
}

func TestMarkdownPrintMessage_EmptyString(t *testing.T) {
	s := ""
	var bb bytes.Buffer
	markdownFormatter.SetOutput(&bb)

	markdownFormatter.PrintMessage(s)
	assert.Equal(t, "\n", bb.String())
}

func BenchmarkMarkdownPrintMessage_EmptyString(b *testing.B) {
	b.ReportAllocs()

	s := ""
	var bb bytes.Buffer
	markdownFormatter.SetOutput(&bb)

	for n := 0; n < b.N; n++ {
		markdownFormatter.PrintMessage(s)
	}
}

func TestMarkdownPrintMessage_TextWithoutNewline(t *testing.T) {
	s := "TEST"
	var bb bytes.Buffer
	markdownFormatter.SetOutput(&bb)

	markdownFormatter.PrintMessage(s)
	assert.Equal(t, fmt.Sprintf("%s  \n", s), bb.String())
}

func TestMarkdownPrintMessage_TextWithNewline(t *testing.T) {
	s := "TEST"
	var bb bytes.Buffer
	markdownFormatter.SetOutput(&bb)

	markdownFormatter.PrintMessage(fmt.Sprintf("%s\n", s))
	assert.Equal(t, fmt.Sprintf("%s  \n", s), bb.String())
}

func TestMarkdownPrintHeader_TextWithoutNewline(t *testing.T) {
	s := "TEST"
	var bb bytes.Buffer
	markdownFormatter.SetOutput(&bb)

	markdownFormatter.PrintHeader(s)

	assert.Equal(t, fmt.Sprintf("\n# %s  \n", s), bb.String())
}

func BenchmarkMarkdownPrintHeader_EmptyString(b *testing.B) {
	b.ReportAllocs()

	s := ""
	var bb bytes.Buffer
	markdownFormatter.SetOutput(&bb)

	for n := 0; n < b.N; n++ {
		markdownFormatter.PrintHeader(s)
	}
}
