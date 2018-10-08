package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToBytes_EmptyString(t *testing.T) {
	var a []byte
	b := convertStringToBytes("")
	assert.IsType(t, b, a)
	assert.Len(t, b, 1)
}

func TestConvertStringToBytes_NewlineOnly(t *testing.T) {
	var a []byte
	b := convertStringToBytes("\n")
	assert.IsType(t, b, a)
	assert.Len(t, b, 1)
}

func TestConvertStringToBytes_TextWithoutNewline(t *testing.T) {
	var a []byte
	b := convertStringToBytes("TEST")
	assert.IsType(t, b, a)
	assert.Len(t, b, 5)
}

func TestConvertStringToBytes_TextWithNewline(t *testing.T) {
	var a []byte
	b := convertStringToBytes("TEST\n")
	assert.IsType(t, b, a)
	assert.Len(t, b, 5)
}
