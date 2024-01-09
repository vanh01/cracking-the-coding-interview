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

func TestURLify(t *testing.T) {
	testCases := []struct {
		S      string
		L      int
		Result string
	}{
		{
			S:      "Mr John Smith     ",
			L:      13,
			Result: "Mr%20John%20Smith",
		},
		{
			S:      "Mr John Smith S      ",
			L:      15,
			Result: "Mr%20John%20Smith%20S",
		},
		{
			S:      "a",
			L:      1,
			Result: "a",
		},
		{
			S:      "",
			L:      0,
			Result: "",
		},
		{
			S:      " a  ",
			L:      2,
			Result: "%20a",
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.URLify(testCase.S, testCase.L)
		require.Equal(t, testCase.Result, result)
	}
}

func TestPalindromePermutation(t *testing.T) {
	testCases := []struct {
		S      string
		Result bool
	}{
		{
			S:      "abcd",
			Result: false,
		},
		{
			S:      "aa",
			Result: true,
		},
		{
			S:      "",
			Result: true,
		},
		{
			S:      "a",
			Result: true,
		},
		{
			S:      "avcavc ",
			Result: true,
		},
		{
			S:      "abcdabc ",
			Result: false,
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.PalindromePermutation(testCase.S)
		require.Equal(t, testCase.Result, result)
	}
}

func TestOneWay(t *testing.T) {
	testCases := []struct {
		S1     string
		S2     string
		Result bool
	}{
		{
			S1:     "abc",
			S2:     "acc",
			Result: true,
		},
		{
			S1:     "abc",
			S2:     "accd",
			Result: false,
		},
		{
			S1:     "abc",
			S2:     "abcf",
			Result: true,
		},
		{
			S1:     "abc",
			S2:     "ab",
			Result: true,
		},
		{
			S1:     "abc",
			S2:     "a",
			Result: false,
		},
		{
			S1:     "abc",
			S2:     "abcde",
			Result: false,
		},
		{
			S1:     "",
			S2:     "",
			Result: true,
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.OneAway(testCase.S1, testCase.S2)
		require.Equal(t, testCase.Result, result)
	}
}
