package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testURLFile = "_TestData/urls.txt"

func TestReadURLFile(t *testing.T) {
	urls, errs := readURLFile(testURLFile)

	assert.Equal(t, 2, len(errs))
	assert.Equal(t, "google.com", errs[0].err.(*url.Error).URL)
	assert.Equal(t, "example.com", errs[1].err.(*url.Error).URL)

	assert.Equal(t, 15, len(urls))
}

func BenchmarkReadURLFile(b *testing.B) {
	b.ReportAllocs()

	Report.SetFormatter("none")

	readURLFile(testURLFile)
}
