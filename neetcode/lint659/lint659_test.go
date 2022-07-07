package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLint659(t *testing.T) {
	testcases := []struct {
		input []string
	}{
		{
			input: []string{"lint", "code", "love", "you"},
		},
		{
			input: []string{"we", "say", ":", "yes"},
		},
	}

	for i, tc := range testcases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			fmt.Printf("input1: %v\n", tc.input)

			act := decode(encode(tc.input))
			assert.Equal(t, tc.input, act)
		})
	}
}
