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

func TestStringCompression(t *testing.T) {
	testCases := []struct {
		S      string
		Result string
	}{
		{
			S:      "abcd",
			Result: "abcd",
		},
		{
			S:      "aaaa",
			Result: "a4",
		},
		{
			S:      "aabcccccaaa",
			Result: "a2b1c5a3",
		},
		{
			S:      "",
			Result: "",
		},
		{
			S:      "a",
			Result: "a",
		},
		{
			S:      "aaaab",
			Result: "a4b1",
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.StringCompression(testCase.S)
		require.Equal(t, testCase.Result, result)
	}
}

func TestRotateMatrix(t *testing.T) {
	testCases := []struct {
		matrix    [][]int
		n         int
		direction int
		result    [][]int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			n:         3,
			direction: 1,
			result: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			n:         3,
			direction: -1,
			result: [][]int{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.RotateMatrix(testCase.matrix, testCase.n, testCase.direction)
		for i := 0; i < testCase.n; i++ {

			for j := 0; j < testCase.n; j++ {
				require.Equal(t, testCase.result[i][j], result[i][j])
			}
		}
	}
}

func TestZeroMatrix(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		m      int
		n      int
		result [][]int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			m: 3,
			n: 3,
			result: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			matrix: [][]int{
				{1, 0, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			m: 3,
			n: 3,
			result: [][]int{
				{0, 0, 0},
				{4, 0, 6},
				{7, 0, 9},
			},
		},
		{
			matrix: [][]int{
				{1, 0, 0, 1},
				{4, 5, 6, 3},
				{7, 8, 9, 9},
			},
			m: 3,
			n: 4,
			result: [][]int{
				{0, 0, 0, 0},
				{4, 0, 0, 3},
				{7, 0, 0, 9},
			},
		},
		{
			matrix: [][]int{
				{1, 0, 0, 1, 3, 4, 8},
				{4, 5, 6, 3, 2, 0, 4},
				{7, 8, 9, 9, 6, 7, 8},
			},
			m: 3,
			n: 7,
			result: [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{7, 0, 0, 9, 6, 0, 8},
			},
		},
	}

	for _, testCase := range testCases {
		result := arraysnstrings.ZeroMatrix(testCase.matrix, testCase.m, testCase.n)
		for i := 0; i < testCase.m; i++ {
			for j := 0; j < testCase.n; j++ {
				require.Equal(t, testCase.result[i][j], result[i][j])
			}
		}
	}
}
