package report

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Formatter Formatter Interface
type Formatter interface {
	SetOutput(output io.Writer)
	PrintMessage(msg string)
	PrintHeader(msg string)
	PrintSubHeader(msg string)
	PrintSeparator()
	PrintError(msg error)
	PrintOK(msg string)
	PrintCode(msg string)
	PrintList([]string)
}

// Reporter Reporter Struct
type Reporter struct {
	Formatter Formatter
}

// New Report
func New() *Reporter {
	return &Reporter{
		Formatter: &ConsoleFormatter{
			Out: os.Stdout,
		},
	}
}

// SetFormatter Set Formatter
func (r *Reporter) SetFormatter(formatType string) {
	fmtType := strings.ToLower(formatType)
	if fmtType == "none" {
		r.Formatter = &ConsoleFormatter{
			Out: ioutil.Discard,
		}
	} else if fmtType == "markdown" {
		r.Formatter = &MarkdownFormatter{
			Out: os.Stdout,
		}
	}
}

// SetOutputFile Set OutputFile
func (r *Reporter) SetOutputFile(filename *string) {
	if *filename == "" {
		return
	}

	outFile, err := os.OpenFile(*filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err == nil {
		r.setOutput(outFile)
	} else {
		r.Formatter.PrintError(errors.New("Failed to create output file, using default stdout"))
	}
}

func (r *Reporter) setOutput(output io.Writer) {
	r.Formatter.SetOutput(output)
}

// PrintMessage Print Message
func (r *Reporter) PrintMessage(msg string) {
	r.Formatter.PrintMessage(msg)
}

// PrintMessagef Print Message Formatted
func (r *Reporter) PrintMessagef(format string, args ...interface{}) {
	r.Formatter.PrintMessage(fmt.Sprintf(format, args...))
}

// PrintHeader Print Header
func (r *Reporter) PrintHeader(msg string) {
	r.Formatter.PrintHeader(msg)
}

// PrintHeaderf Print Header Formatted
func (r *Reporter) PrintHeaderf(format string, args ...interface{}) {
	r.Formatter.PrintHeader(fmt.Sprintf(format, args...))
}

// PrintSubHeader Print Sub-Header
func (r *Reporter) PrintSubHeader(msg string) {
	r.Formatter.PrintSubHeader(msg)
}

// PrintSubHeaderf Print Header Formatted
func (r *Reporter) PrintSubHeaderf(format string, args ...interface{}) {
	r.Formatter.PrintSubHeader(fmt.Sprintf(format, args...))
}

// PrintSeparator Print Separator
func (r *Reporter) PrintSeparator() {
	r.Formatter.PrintSeparator()
}

// PrintError Print Error
func (r *Reporter) PrintError(msg error) {
	r.Formatter.PrintError(msg)
}

// PrintOK Print OK
func (r *Reporter) PrintOK(msg string) {
	r.Formatter.PrintOK(msg)
}

// PrintOKf Print OK Formatted
func (r *Reporter) PrintOKf(format string, args ...interface{}) {
	r.Formatter.PrintOK(fmt.Sprintf(format, args...))
}

// PrintCode Print Code
func (r *Reporter) PrintCode(msg string) {
	r.Formatter.PrintCode(msg)
}

// PrintCodef Print Code Formatted
func (r *Reporter) PrintCodef(format string, args ...interface{}) {
	r.Formatter.PrintCode(fmt.Sprintf(format, args...))
}

// PrintList Print List
func (r *Reporter) PrintList(list []string) {
	r.Formatter.PrintList(list)
}
