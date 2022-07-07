package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP128(t *testing.T) {
	testcases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			input:    []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input)
			assert.Equal(t, tc.expected, longestConsecutive(tc.input))
		})
	}
}
