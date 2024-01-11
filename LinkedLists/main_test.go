package linkedlists_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	linkedlists "github.com/vanh01/cracking-the-coding-interview/LinkedLists"
)

func TestAppend(t *testing.T) {
	testCases := []struct {
		node    *linkedlists.Node[int]
		newNode *linkedlists.Node[int]
		result  []int
	}{
		{
			node:    &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2}},
			newNode: &linkedlists.Node[int]{Data: 3},
			result:  []int{1, 2, 3},
		},
		{
			node:    &linkedlists.Node[int]{Data: 1},
			newNode: &linkedlists.Node[int]{Data: 2},
			result:  []int{1, 2},
		},
	}

	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		linkedList.Append(testCase.newNode)
		require.Equal(t, testCase.result, linkedList.Root.ToArray())
	}
}

func TestInsertAt(t *testing.T) {
	testCases := []struct {
		node    *linkedlists.Node[int]
		newNode *linkedlists.Node[int]
		index   int
		result  []int
	}{
		{
			node:    &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2}},
			newNode: &linkedlists.Node[int]{Data: 3},
			index:   2,
			result:  []int{1, 2, 3},
		},
		{
			node:    &linkedlists.Node[int]{Data: 1},
			newNode: &linkedlists.Node[int]{Data: 2},
			index:   0,
			result:  []int{2, 1},
		},
		{
			node:    &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2, Next: &linkedlists.Node[int]{Data: 3}}},
			newNode: &linkedlists.Node[int]{Data: 9},
			index:   -1,
			result:  []int{9, 1, 2, 3},
		},
		{
			node:    &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2, Next: &linkedlists.Node[int]{Data: 3}}},
			newNode: &linkedlists.Node[int]{Data: 9},
			index:   4,
			result:  []int{1, 2, 3, 9},
		},
		{
			node:    &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2, Next: &linkedlists.Node[int]{Data: 3}}},
			newNode: &linkedlists.Node[int]{Data: 9},
			index:   5,
			result:  []int{1, 2, 3, 9},
		},
		{
			node:    &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2, Next: &linkedlists.Node[int]{Data: 3}}},
			newNode: &linkedlists.Node[int]{Data: 9},
			index:   2,
			result:  []int{1, 2, 9, 3},
		},
	}

	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		linkedList.InsertAt(testCase.newNode, testCase.index)
		require.Equal(t, testCase.result, linkedList.Root.ToArray())
	}
}

func TestDeleteAt(t *testing.T) {
	testCases := []struct {
		node   *linkedlists.Node[int]
		index  int
		result []int
	}{
		{
			node:   &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2}},
			index:  1,
			result: []int{1},
		},
		{
			node:   &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2}},
			index:  0,
			result: []int{2},
		},
		{
			node:   &linkedlists.Node[int]{Data: 1},
			index:  0,
			result: []int{},
		},
		{
			node:   &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2, Next: &linkedlists.Node[int]{Data: 3}}},
			index:  2,
			result: []int{1, 2},
		},
		{
			node:   &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2, Next: &linkedlists.Node[int]{Data: 3}}},
			index:  0,
			result: []int{2, 3},
		},
		{
			node:   &linkedlists.Node[int]{Data: 1, Next: &linkedlists.Node[int]{Data: 2, Next: &linkedlists.Node[int]{Data: 3}}},
			index:  1,
			result: []int{1, 3},
		},
	}

	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		linkedList.DeleteAt(testCase.index)
		require.Equal(t, testCase.result, linkedList.Root.ToArray())
	}
}

func TestRemoveDups(t *testing.T) {
	testCases := []struct {
		node   *linkedlists.Node[int]
		result []int
	}{
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			result: []int{1, 2, 3},
		},
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{
						Data: 3,
						Next: &linkedlists.Node[int]{Data: 3},
					},
				},
			},
			result: []int{1, 2, 3},
		},
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 1,
					Next: &linkedlists.Node[int]{
						Data: 1,
						Next: &linkedlists.Node[int]{Data: 1},
					},
				},
			},
			result: []int{1},
		},
	}

	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		result := linkedlists.RemoveDups[int](linkedList)
		require.Equal(t, testCase.result, result.Root.ToArray())
	}
}
