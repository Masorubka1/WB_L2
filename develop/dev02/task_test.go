package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string")},
		{"", "", nil},
	}

	for _, tt := range tests {
		actual, err := UnpackString(tt.input)
		assert.Equal(t, tt.expected, actual)
		assert.Equal(t, tt.err, err)
	}
}
