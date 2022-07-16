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
			input1:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			input1:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			assert.Equal(t, tc.expected, maxProfit(tc.input1))
		})
	}
}
