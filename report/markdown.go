package report

import (
	"fmt"
	"io"
	"sync"
)

// MarkdownFormatter Markdown Formatter
type MarkdownFormatter struct {
	Out io.Writer
	mu  sync.Mutex
}

// SetOutput Set Output Writer
func (f *MarkdownFormatter) SetOutput(output io.Writer) {
	f.mu.Lock()
	f.Out = output
	defer f.mu.Unlock()
}

// PrintMessage Print Markdown Message
func (f *MarkdownFormatter) PrintMessage(msg string) {
	padMsg := fmt.Sprintf("%s  ", msg)
	b := convertStringToBytes(padMsg)
	f.mu.Lock()
	f.Out.Write(b)
	defer f.mu.Unlock()
}

// PrintHeader Print Markdown Header Message
func (f *MarkdownFormatter) PrintHeader(msg string) {
	f.PrintMessage(fmt.Sprintf("# %s", msg))
}

// PrintSubHeader Print Markdown Sub-Header Message
func (f *MarkdownFormatter) PrintSubHeader(msg string) {
	f.PrintMessage(fmt.Sprintf("## %s", msg))
}

// PrintSeparator Print Markdown Separator
func (f *MarkdownFormatter) PrintSeparator() {
	f.PrintMessage("---")
}

// PrintError Print Markdown Error
func (f *MarkdownFormatter) PrintError(msg error) {
	f.PrintMessage(fmt.Sprintf("> ERROR: %s", msg))
}

// PrintOK Print Markdown Error
func (f *MarkdownFormatter) PrintOK(msg string) {
	f.PrintMessage(fmt.Sprintf("> OK: %s", msg))
}

// PrintCode Print Markdown Code
func (f *MarkdownFormatter) PrintCode(msg string) {
	f.PrintMessage("~~~")
	f.PrintMessage(msg)
	f.PrintMessage("~~~")
}

// PrintList Print Markdown List
func (f *MarkdownFormatter) PrintList(list []string) {
	f.PrintMessage("")
	for _, s := range list {
		f.PrintMessage(fmt.Sprintf("- %s", s))
	}
	f.PrintMessage("")
}
