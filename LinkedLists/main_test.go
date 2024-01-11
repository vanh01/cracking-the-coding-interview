package linkedlists_test

import (
	"fmt"
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

func randomNode(l int) *linkedlists.Node[int] {
	root := &linkedlists.Node[int]{
		Data: 1,
	}
	node := root
	for i := 1; i < l; i++ {
		node.Next = &linkedlists.Node[int]{
			Data: i + 1,
		}
		node = node.Next
	}
	return root
}

func TestDeepCopy(t *testing.T) {
	testCases := []*linkedlists.Node[int]{
		{},
		randomNode(1),
		randomNode(10),
		randomNode(5),
	}
	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase,
		}
		result := linkedList.DeepCopy()
		root1 := linkedList.Root
		root2 := result.Root

		for root1 != nil {
			fmt.Println("Begining traveling all node")
			fmt.Printf(">>> node1: %v, node2: %v\n", root1, root2)
			require.Equal(t, root1.Data, root2.Data)
			require.NotSame(t, root1, root2)
			root1 = root1.Next
			root2 = root2.Next
		}
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
