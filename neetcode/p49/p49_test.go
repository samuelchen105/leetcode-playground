package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP49(t *testing.T) {
	testcases := []struct {
		input1   []string
		expected [][]string
	}{
		{
			input1:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			input1:   []string{""},
			expected: [][]string{{""}},
		},
		{
			input1:   []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input1)
			// may fail to pass, ElementsMatch does not support slice
			assert.ElementsMatch(t, tc.expected, groupAnagrams(tc.input1))
		})
	}
}
