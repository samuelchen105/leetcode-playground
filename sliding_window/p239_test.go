package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP239(t *testing.T) {
	testcases := []struct {
		input1   []int
		input2   int
		expected []int
	}{
		{
			input1:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			input2:   3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			input1:   []int{1},
			input2:   1,
			expected: []int{1},
		},
		{
			input1:   []int{1, -1},
			input2:   1,
			expected: []int{1, -1},
		},
		{
			input1:   []int{9, 11},
			input2:   2,
			expected: []int{11},
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			fmt.Printf("input2: %v\n", tc.input2)
			assert.Equal(t, tc.expected, maxSlidingWindow(tc.input1, tc.input2))
		})
	}
}
