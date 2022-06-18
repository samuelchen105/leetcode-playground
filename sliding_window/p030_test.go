package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP030(t *testing.T) {
	testcases := []struct {
		input1   string
		input2   []string
		expected []int
	}{
		{
			input1:   "barfoothefoobarman",
			input2:   []string{"foo", "bar"},
			expected: []int{0, 9},
		},
		{
			input1:   "wordgoodgoodgoodbestword",
			input2:   []string{"word", "good", "best", "word"},
			expected: []int{},
		},
		{
			input1:   "barfoofoobarthefoobarman",
			input2:   []string{"bar", "foo", "the"},
			expected: []int{6, 9, 12},
		},
		{
			input1:   "wordgoodgoodgoodbestword",
			input2:   []string{"word", "good", "best", "good"},
			expected: []int{8},
		},
		{
			input1:   "lingmindraboofooowingdingbarrwingmonkeypoundcake",
			input2:   []string{"fooo", "barr", "wing", "ding", "wing"},
			expected: []int{13},
		},
		{
			input1:   "aaaaaaaaaaaaaa",
			input2:   []string{"aa", "aa"},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %s\n", tc.input1)
			fmt.Printf("input2: %v\n", tc.input2)
			assert.ElementsMatch(t, tc.expected, findSubstring(tc.input1, tc.input2))
		})
	}
}
