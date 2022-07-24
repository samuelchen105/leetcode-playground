package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) {
	testcases := []struct {
		input1   string
		expected int
	}{
		{
			input1:   "abcabcbb",
			expected: 3,
		},
		{
			input1:   "bbbbb",
			expected: 1,
		},
		{
			input1:   "pwwkew",
			expected: 3,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			assert.Equal(t, tc.expected, lengthOfLongestSubstring(tc.input1))
		})
	}
}
