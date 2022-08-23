package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) {
	testcases := []struct {
		input1   string
		expected bool
	}{
		{
			input1:   "()",
			expected: true,
		},
		{
			input1:   "()[]{}",
			expected: true,
		},
		{
			input1:   "(]",
			expected: false,
		},
		{
			input1:   "([)]",
			expected: false,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			assert.Equal(t, tc.expected, isValid(tc.input1))
		})
	}
}
