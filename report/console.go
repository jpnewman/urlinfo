package report

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// ConsoleWidth Width of headers
var ConsoleWidth = 80

// ConsoleFormatter Console Formatter
type ConsoleFormatter struct {
	Out io.Writer
	mu  sync.Mutex
}

// SetOutput Set Output Writer
func (f *ConsoleFormatter) SetOutput(output io.Writer) {
	f.mu.Lock()
	f.Out = output
	defer f.mu.Unlock()
}

// PrintMessage Print Console Message
func (f *ConsoleFormatter) PrintMessage(msg string) {
	b := convertStringToBytes(msg)
	f.mu.Lock()
	f.Out.Write(b)
	defer f.mu.Unlock()
}

// TODO: Review need for lock around color if output is a file, etc.
// printMessage Print Console Formatted Message
func (f *ConsoleFormatter) printFormattedMessage(msg string, overline string, underline string, width int, c *color.Color) {
	f.mu.Lock()
	color.Output = f.Out

	c.Println(strings.Repeat(overline, ConsoleWidth))
	c.Println(msg)
	c.Println(strings.Repeat(underline, ConsoleWidth))

	defer f.mu.Unlock()
}

// PrintHeader Prints Console Header
func (f *ConsoleFormatter) PrintHeader(msg string) {
	c := color.New(color.FgCyan, color.Bold)
	f.printFormattedMessage(msg, "=", "=", ConsoleWidth, c)
}

// PrintSubHeader Prints Console Sub-Header
func (f *ConsoleFormatter) PrintSubHeader(msg string) {
	c := color.New(color.FgBlue)
	f.printFormattedMessage(msg, "-", "-", ConsoleWidth, c)
}

// PrintSeparator Print Console Separator
func (f *ConsoleFormatter) PrintSeparator() {
	color.Magenta(fmt.Sprintf(strings.Repeat("-", ConsoleWidth)))
}

// PrintError Print Console Error
func (f *ConsoleFormatter) PrintError(msg error) {
	c := color.New(color.FgRed)
	c.Printf("ERROR: %s\n", msg)
}

// PrintOK Print Console OK
func (f *ConsoleFormatter) PrintOK(msg string) {
	c := color.New(color.FgGreen)
	c.Printf("OK: %s\n", msg)
}

// PrintCode Print Console Code
func (f *ConsoleFormatter) PrintCode(msg string) {
	c := color.New(color.FgMagenta)

	scanner := bufio.NewScanner(strings.NewReader(msg))
	for scanner.Scan() {
		c.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		f.PrintError(err)
	}
}

// PrintList Print Console List
func (f *ConsoleFormatter) PrintList(list []string) {
	for _, s := range list {
		f.PrintMessage(fmt.Sprintf("- %s", s))
	}
}
