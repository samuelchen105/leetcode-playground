package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP128(t *testing.T) {
	testcases := []struct {
		input1   []int
		expected [][]int
	}{
		{
			input1:   []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input1:   []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			input1:   []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			assert.Equal(t, tc.expected, threeSum(tc.input1))
		})
	}
}
