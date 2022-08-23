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
		expected int
	}{
		{
			input1:   []int{4, 5, 6, 7, 0, 1, 2},
			input2:   0,
			expected: 4,
		},
		{
			input1:   []int{4, 5, 6, 7, 0, 1, 2},
			input2:   3,
			expected: -1,
		},
		{
			input1:   []int{1},
			input2:   0,
			expected: -1,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			assert.Equal(t, tc.expected, search(tc.input1, tc.input2))
		})
	}
}
