package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_centerString(t *testing.T) {
	assert.Equal(t, "   8  ", centerString("8", 6))
	assert.Equal(t, "  16  ", centerString("16", 6))
	assert.Equal(t, "  128 ", centerString("128", 6))
	assert.Equal(t, " 1024 ", centerString("1024", 6))
}
