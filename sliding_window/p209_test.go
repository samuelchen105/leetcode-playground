package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP209(t *testing.T) {
	testcases := []struct {
		input1   int
		input2   []int
		expected int
	}{
		{
			input1:   7,
			input2:   []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			input1:   4,
			input2:   []int{1, 4, 4},
			expected: 1,
		},
		{
			input1:   11,
			input2:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			fmt.Printf("input2: %v\n", tc.input2)
			assert.Equal(t, tc.expected, minSubArrayLen(tc.input1, tc.input2))
		})
	}
}
