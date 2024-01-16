package linkedlists_test

import (
	"fmt"
	"math/rand"
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
	if l <= 0 {
		return nil
	}
	root := &linkedlists.Node[int]{
		Data: 1,
	}
	node := root
	for i := 1; i < l; i++ {
		node.Next = &linkedlists.Node[int]{
			Data: rand.Intn(l),
		}
		node = node.Next
	}
	return root
}

func TestDeepCopy(t *testing.T) {
	testCases := []*linkedlists.Node[int]{
		nil,
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

		fmt.Println("Begining traveling all node")
		for root1 != nil {
			fmt.Printf(">>> node1: %v, node2: %v\n", root1, root2)
			require.Equal(t, root1.Data, root2.Data)
			require.NotSame(t, root1, root2)
			root1 = root1.Next
			root2 = root2.Next
		}
	}
}

func TestGetLength(t *testing.T) {
	testCases := []struct {
		node   *linkedlists.Node[int]
		result int
	}{
		{node: randomNode(0), result: 0},
		{node: randomNode(1), result: 1},
		{node: randomNode(5), result: 5},
		{node: randomNode(10), result: 10},
	}
	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		result := linkedList.Length()
		require.Equal(t, testCase.result, result)
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

func TestReturnKthToLast(t *testing.T) {
	testCases := []struct {
		node   *linkedlists.Node[int]
		k      int
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
			k:      1,
			result: []int{2, 3},
		},
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			k:      2,
			result: []int{3},
		},
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{
						Data: 3,
						Next: &linkedlists.Node[int]{
							Data: 4,
							Next: &linkedlists.Node[int]{Data: 5},
						},
					},
				},
			},
			k:      2,
			result: []int{3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		result := linkedlists.ReturnKthToLast[int](linkedList, testCase.k)
		require.Equal(t, testCase.result, result.Root.ToArray())
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	testCases := []struct {
		node   *linkedlists.Node[int]
		result []int
	}{
		{node: randomNode(10)},
		{node: randomNode(9)},
		{node: randomNode(3)},
		{node: randomNode(2)},
		{node: randomNode(1)},
	}
	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		expected := testCase.node.ToArray()
		middle := (linkedList.Length() - 1) / 2
		if middle > 0 {
			expected = append(expected[:middle], expected[middle+1:]...)
		}
		linkedlists.DeleteMiddleNode[int](linkedList)
		require.Equal(t, expected, linkedList.Root.ToArray())
	}
}

func TestPartition(t *testing.T) {

	check := func(root *linkedlists.Node[int], partition int) bool {
		flag := false
		node := root

		for node != nil {
			if flag && node.Data < partition {
				return false
			}
			if node.Data >= partition {
				flag = true
			}
			node = node.Next
		}
		return true
	}

	testCases := []struct {
		node      *linkedlists.Node[int]
		partition int
		check     func(*linkedlists.Node[int], int) bool
	}{
		{node: randomNode(10), partition: rand.Intn(8), check: check},
		{node: randomNode(30), partition: rand.Intn(20), check: check},
	}

	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		fmt.Println(linkedList.Root.ToArray(), testCase.partition)
		linkedlists.Partition[int](linkedList, testCase.partition)
		fmt.Println(linkedList.Root.ToArray())
		require.True(t, testCase.check(linkedList.Root, testCase.partition))
	}
}

func TestSumLists(t *testing.T) {
	testCases := []struct {
		num1   *linkedlists.Node[int]
		num2   *linkedlists.Node[int]
		result []int
	}{
		{
			num1: &linkedlists.Node[int]{
				Data: 0,
			},
			num2: &linkedlists.Node[int]{
				Data: 0,
			},
			result: []int{0},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 1,
			},
			num2: &linkedlists.Node[int]{
				Data: 1,
			},
			result: []int{2},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			num2: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			result: []int{2, 4, 6},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 0,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			num2: &linkedlists.Node[int]{
				Data: 0,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			result: []int{0, 0, 6},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 0,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 7},
				},
			},
			num2: &linkedlists.Node[int]{
				Data: 0,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			result: []int{0, 0, 0, 1},
		},
	}
	for _, testCase := range testCases {
		linkedList1 := &linkedlists.LinkedList[int]{
			Root: testCase.num1,
		}
		linkedList2 := &linkedlists.LinkedList[int]{
			Root: testCase.num2,
		}
		result := linkedlists.SumLists[int](linkedList1, linkedList2)
		require.Equal(t, testCase.result, result.Root.ToArray())
	}
}

func TestSumListsV2(t *testing.T) {
	testCases := []struct {
		num1   *linkedlists.Node[int]
		num2   *linkedlists.Node[int]
		result []int
	}{
		{
			num1: &linkedlists.Node[int]{
				Data: 0,
			},
			num2: &linkedlists.Node[int]{
				Data: 0,
			},
			result: []int{0},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 1,
			},
			num2: &linkedlists.Node[int]{
				Data: 1,
			},
			result: []int{2},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			num2: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 3},
				},
			},
			result: []int{2, 4, 6},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 3,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 0},
				},
			},
			num2: &linkedlists.Node[int]{
				Data: 3,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 0},
				},
			},
			result: []int{6, 0, 0},
		},
		{
			num1: &linkedlists.Node[int]{
				Data: 7,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 0},
				},
			},
			num2: &linkedlists.Node[int]{
				Data: 3,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 0},
				},
			},
			result: []int{1, 0, 0, 0},
		},
	}
	for _, testCase := range testCases {
		linkedList1 := &linkedlists.LinkedList[int]{
			Root: testCase.num1,
		}
		linkedList2 := &linkedlists.LinkedList[int]{
			Root: testCase.num2,
		}
		result := linkedlists.SumListsV2[int](linkedList1, linkedList2)
		require.Equal(t, testCase.result, result.Root.ToArray())
	}
}

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		node   *linkedlists.Node[int]
		result bool
	}{
		{
			node:   &linkedlists.Node[int]{Data: 0},
			result: true,
		},
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 1},
				},
			},
			result: true,
		},
		{
			node: &linkedlists.Node[int]{
				Data: 0,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 1},
				},
			},
			result: false,
		},
	}

	for _, testCase := range testCases {
		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		result := linkedlists.Palindrome[int](linkedList)
		require.Equal(t, testCase.result, result)
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		node1        *linkedlists.Node[int]
		node2        *linkedlists.Node[int]
		intersection *linkedlists.Node[int]
	}{
		{
			node1: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 2},
				},
			},
			node2:        &linkedlists.Node[int]{Data: 1},
			intersection: nil,
		},
		{
			node1: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 2},
				},
			},
			node2: &linkedlists.Node[int]{Data: 1},
			intersection: &linkedlists.Node[int]{
				Data: 11,
				Next: &linkedlists.Node[int]{
					Data: 10,
					Next: &linkedlists.Node[int]{Data: 12},
				},
			},
		},
		{
			node1: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 0,
					Next: &linkedlists.Node[int]{Data: 2},
				},
			},
			node2:        &linkedlists.Node[int]{Data: 1},
			intersection: &linkedlists.Node[int]{Data: 11},
		},
	}

	for _, testCase := range testCases {
		linkedList1 := &linkedlists.LinkedList[int]{
			Root: testCase.node1,
		}
		linkedList2 := &linkedlists.LinkedList[int]{
			Root: testCase.node2,
		}
		linkedList1.Append(testCase.intersection)
		linkedList2.Append(testCase.intersection)
		result := linkedlists.Intersection[int](linkedList1, linkedList2)
		require.Equal(t, testCase.intersection, result)
	}
}

func TestLoopDetection(t *testing.T) {
	testCases := []struct {
		node     *linkedlists.Node[int]
		loopNode *linkedlists.Node[int]
	}{
		{},
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 33},
				},
			},
			loopNode: &linkedlists.Node[int]{
				Data: 4,
				Next: &linkedlists.Node[int]{
					Data: 5,
					Next: &linkedlists.Node[int]{Data: 6},
				},
			},
		},
		{
			node: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 33},
				},
			},
		},
		{
			loopNode: &linkedlists.Node[int]{
				Data: 1,
				Next: &linkedlists.Node[int]{
					Data: 2,
					Next: &linkedlists.Node[int]{Data: 33},
				},
			},
		},
	}

	for _, testCase := range testCases {
		node := testCase.node
		if testCase.node == nil && testCase.loopNode == nil {
			continue
		}
		if node != nil {
			for node.Next != nil {
				node = node.Next
			}
			node.Next = testCase.loopNode
		}

		if testCase.loopNode != nil {
			node = testCase.loopNode
		}
		for node.Next != nil {
			node = node.Next
		}
		node.Next = testCase.loopNode

		linkedList := &linkedlists.LinkedList[int]{
			Root: testCase.node,
		}
		if linkedList.Root == nil {
			linkedList.Root = testCase.loopNode
		}

		result := linkedlists.LoopDetection[int](linkedList)
		require.Equal(t, testCase.loopNode, result)
	}
}
