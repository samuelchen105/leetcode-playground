package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP424(t *testing.T) {
	testcases := []struct {
		input1   string
		input2   int
		expected int
	}{
		{
			input1:   "ABAB",
			input2:   2,
			expected: 4,
		},
		{
			input1:   "AABABBA",
			input2:   1,
			expected: 4,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			fmt.Printf("input2: %v\n", tc.input2)
			assert.Equal(t, tc.expected, characterReplacement(tc.input1, tc.input2))
		})
	}
}
