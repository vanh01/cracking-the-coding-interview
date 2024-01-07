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
