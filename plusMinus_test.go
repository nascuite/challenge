package main

import "testing"

type testCase struct {
	input  int
	output string
}

func TestPlusMinus(t *testing.T) {
	testCases := []testCase{
		{
			input:  35132,
			output: "-++-",
		},
		{
			input:  351,
			output: "unreal",
		},
	}

	for i, c := range testCases {
		res := PlusMinus(c.input)
		if res != c.output {
			t.Errorf("[%d] expected %s, got %s", i, c.output, res)
		}
	}
}
