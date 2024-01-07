package arraysnstrings_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	arraysnstrings "arraynstring/ArraysnStrings"
)

func TestIsUnique(t *testing.T) {
	testCases := []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "",
			Output: true,
		},
		{
			Input:  "abc",
			Output: true,
		},
		{
			Input:  "abcddssdsdds",
			Output: false,
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.IsUnique(testCase.Input)
		require.Equal(t, testCase.Output, result)
	}
}

func TestCheckPermutation(t *testing.T) {
	testCases := []struct {
		S1     string
		S2     string
		Result bool
	}{
		{
			S1:     "123",
			S2:     "123",
			Result: true,
		},
		{
			S1:     "123",
			S2:     "231",
			Result: true,
		},
		{
			S1:     "123",
			S2:     "213",
			Result: true,
		},
		{
			S1:     "1231",
			S2:     "213",
			Result: false,
		},
		{
			S1:     "",
			S2:     "213",
			Result: false,
		},
		{
			S1:     "",
			S2:     "",
			Result: true,
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.CheckPermutation(testCase.S1, testCase.S2)
		require.Equal(t, testCase.Result, result)
	}
}
