package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) {
	testcases := []struct {
		input1   []int
		expected int
	}{
		{
			input1:   []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			input1:   []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			input1:   []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			input1:   []int{1},
			expected: 1,
		},
		{
			input1:   []int{3, 1, 2},
			expected: 1,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			assert.Equal(t, tc.expected, findMin(tc.input1))
		})
	}
}
