package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) {
	testcases := []struct {
		input1   []int
		input2   int
		expected []int
	}{
		{
			input1:   []int{2, 7, 11, 15},
			input2:   9,
			expected: []int{1, 2},
		},
		{
			input1:   []int{2, 3, 4},
			input2:   6,
			expected: []int{1, 3},
		},
		{
			input1:   []int{-1, 0},
			input2:   -1,
			expected: []int{1, 2},
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			fmt.Printf("input2: %v\n", tc.input2)
			assert.Equal(t, tc.expected, twoSum(tc.input1, tc.input2))
		})
	}
}
