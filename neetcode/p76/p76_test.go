package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) {
	testcases := []struct {
		input1   string
		input2   string
		expected string
	}{
		{
			input1:   "ADOBECODEBANC",
			input2:   "ABC",
			expected: "BANC",
		},
		{
			input1:   "a",
			input2:   "a",
			expected: "a",
		},
		{
			input1:   "a",
			input2:   "aa",
			expected: "",
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			assert.Equal(t, tc.expected, minWindow(tc.input1, tc.input2))
		})
	}
}
