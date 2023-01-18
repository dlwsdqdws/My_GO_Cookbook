package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func Hello() string {
// 	return "hi"
// }

func TestHello(t *testing.T) {
	output := Hello()
	expectedOutput := "hello"
	// if output != expectedOutput {
	// 	t.Errorf("Expected %s do not match actual %s", expectedOutput, output)
	// }
	assert.Equal(t, expectOutput, output)
}
