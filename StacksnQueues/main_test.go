package stacksnqueues_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	stacksnqueues "github.com/vanh01/cracking-the-coding-interview/StacksnQueues"
)

func TestPush(t *testing.T) {
	testCases := []struct {
		ops []struct {
			typ   stacksnqueues.StackType
			value int
		}
		result []int
	}{
		{
			ops: []struct {
				typ   stacksnqueues.StackType
				value int
			}{
				{stacksnqueues.SECOND, 1},
				{stacksnqueues.FIRST, 2},
				{stacksnqueues.SECOND, 0},
				{stacksnqueues.THRID, 1},
				{stacksnqueues.FIRST, 3},
				{stacksnqueues.THRID, 9},
			},
			result: []int{2, 3, 1, 0, 1, 9},
		},
	}

	stack := stacksnqueues.NewThreeStack()
	for _, testCase := range testCases {
		for _, op := range testCase.ops {
			stack.Push(op.typ, op.value)
		}
		require.Equal(t, testCase.result, stack.ToArray())
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		stack stacksnqueues.ThreeStack
		ops   []struct {
			typ   stacksnqueues.StackType
			value int
		}
		result []int
	}{
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 6, 0, 0),
			ops: []struct {
				typ   stacksnqueues.StackType
				value int
			}{
				{stacksnqueues.FIRST, 6},
				{stacksnqueues.FIRST, 5},
				{stacksnqueues.FIRST, 4},
			},
			result: []int{1, 2, 3},
		},
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 3, 2, 1),
			ops: []struct {
				typ   stacksnqueues.StackType
				value int
			}{
				{stacksnqueues.FIRST, 3},
				{stacksnqueues.FIRST, 2},
				{stacksnqueues.FIRST, 1},
			},
			result: []int{4, 5, 6},
		},
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 3, 2, 1),
			ops: []struct {
				typ   stacksnqueues.StackType
				value int
			}{
				{stacksnqueues.SECOND, 5},
				{stacksnqueues.FIRST, 3},
				{stacksnqueues.FIRST, 2},
			},
			result: []int{1, 4, 6},
		},
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 3, 2, 1),
			ops: []struct {
				typ   stacksnqueues.StackType
				value int
			}{
				{stacksnqueues.SECOND, 5},
				{stacksnqueues.THRID, 6},
				{stacksnqueues.FIRST, 3},
				{stacksnqueues.FIRST, 2},
			},
			result: []int{1, 4},
		},
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 3, 2, 1),
			ops: []struct {
				typ   stacksnqueues.StackType
				value int
			}{
				{stacksnqueues.SECOND, 5},
				{stacksnqueues.THRID, 6},
				{stacksnqueues.FIRST, 3},
				{stacksnqueues.FIRST, 2},
				{stacksnqueues.FIRST, 1},
				{stacksnqueues.SECOND, 4},
			},
			result: []int{},
		},
	}
	for _, testCase := range testCases {
		stack := testCase.stack
		for _, op := range testCase.ops {
			value := stack.Pop(op.typ)
			require.Equal(t, op.value, value)
		}
		require.Equal(t, testCase.result, stack.ToArray())
	}
}

func TestPeek(t *testing.T) {
	testCases := []struct {
		stack stacksnqueues.ThreeStack
		ops   []struct {
			typ   stacksnqueues.StackType
			value int
		}
	}{
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 3, 2, 1),
			ops: []struct {
				typ   stacksnqueues.StackType
				value int
			}{
				{stacksnqueues.FIRST, 3},
				{stacksnqueues.SECOND, 5},
				{stacksnqueues.THRID, 6},
			},
		},
	}

	for _, testCase := range testCases {
		stack := testCase.stack
		for _, op := range testCase.ops {
			value := stack.Peek(op.typ)
			require.Equal(t, op.value, value)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		stack stacksnqueues.ThreeStack
		ops   []struct {
			typ    stacksnqueues.StackType
			result bool
		}
	}{
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 3, 2, 1),
			ops: []struct {
				typ    stacksnqueues.StackType
				result bool
			}{
				{stacksnqueues.FIRST, false},
				{stacksnqueues.SECOND, false},
				{stacksnqueues.THRID, false},
			},
		},
		{
			stack: stacksnqueues.NewThreeStackFull([]int{1, 2, 3, 4, 5, 6}, 6, 0, 0),
			ops: []struct {
				typ    stacksnqueues.StackType
				result bool
			}{
				{stacksnqueues.FIRST, false},
				{stacksnqueues.SECOND, true},
				{stacksnqueues.THRID, true},
			},
		},
	}

	for _, testCase := range testCases {
		stack := testCase.stack
		for _, op := range testCase.ops {
			value := stack.IsEmpty(op.typ)
			require.Equal(t, op.result, value)
		}
	}
}
