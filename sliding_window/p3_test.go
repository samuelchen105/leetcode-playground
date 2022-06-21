package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP3(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "abba",
			expected: 2,
		},
	}

	for _, tc := range testcases {
		name := fmt.Sprintf("\"%s\"", tc.input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, lengthOfLongestSubstring(tc.input))
		})
	}
}
