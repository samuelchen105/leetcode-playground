package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP242(t *testing.T) {
	testcases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{
			input1:   "anagram",
			input2:   "nagaram",
			expected: true,
		},
		{
			input1:   "rat",
			input2:   "car",
			expected: false,
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			fmt.Printf("input2: %v\n", tc.input2)
			assert.Equal(t, tc.expected, isAnagram(tc.input1, tc.input2))
		})
	}
}
