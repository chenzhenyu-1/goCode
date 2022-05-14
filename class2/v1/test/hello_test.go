package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	output := hello(80)
	expectOutput := "hello"
	assert.Equal(t, expectOutput, output)
}
