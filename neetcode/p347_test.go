package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP347(t *testing.T) {
	testcases := []struct {
		input1   []int
		input2   int
		expected []int
	}{
		{
			input1:   []int{1, 1, 1, 2, 2, 3},
			input2:   2,
			expected: []int{1, 2},
		},
		{
			input1:   []int{1},
			input2:   1,
			expected: []int{1},
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			fmt.Printf("input2: %v\n", tc.input2)
			assert.Equal(t, tc.expected, topKFrequent(tc.input1, tc.input2))
		})
	}
}
