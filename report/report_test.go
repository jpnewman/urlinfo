package report

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

var consoleFormatter *ConsoleFormatter
var markdownFormatter *MarkdownFormatter

func setup() {
	fmt.Println("Test Setup...")

	var bbc bytes.Buffer
	consoleWriter := bufio.NewWriter(&bbc)
	consoleFormatter = &ConsoleFormatter{
		Out: consoleWriter,
	}

	var bbm bytes.Buffer
	markdownWriter := bufio.NewWriter(&bbm)

	markdownFormatter = &MarkdownFormatter{
		Out: markdownWriter,
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
