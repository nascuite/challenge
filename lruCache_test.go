package main

import (
	"testing"
)

type testCases struct {
	input  []string
	output string
}

func TestCacheLRU(t *testing.T) {
	testCases := []testCases{
		{
			input:  []string{"A", "B", "A", "C", "A"},
			output: "BCA",
		},
		{
			input:  []string{"A", "A", "A"},
			output: "A",
		},
	}

	for i, testCase := range testCases {
		res := cacheLRU(testCase.input)
		if res != testCase.output {
			t.Errorf("test case num=%d expected: %s, got: %s", i, testCase.input, res)
		}
	}

}
