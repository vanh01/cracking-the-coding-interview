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

func TestStackMin(t *testing.T) {
	testCases := []struct {
		stack stacksnqueues.StackMin
		ops   []struct {
			typ    stacksnqueues.StackOpType
			value  int
			min    int
			result []int
		}
	}{
		{
			stack: stacksnqueues.StackMin{},
			ops: []struct {
				typ    stacksnqueues.StackOpType
				value  int
				min    int
				result []int
			}{
				{
					typ:    stacksnqueues.PUSH,
					value:  1,
					min:    1,
					result: []int{1},
				},
				{
					typ:    stacksnqueues.PUSH,
					value:  2,
					min:    1,
					result: []int{1, 2},
				},
				{
					typ:    stacksnqueues.PUSH,
					value:  0,
					min:    0,
					result: []int{1, 2, 0},
				},
				{
					typ:    stacksnqueues.POP,
					min:    1,
					result: []int{1, 2},
				},
			},
		},
	}

	for _, testCase := range testCases {
		stack := testCase.stack
		for _, op := range testCase.ops {
			switch op.typ {
			case stacksnqueues.PUSH:
				stack.Push(op.value)
				require.Equal(t, op.min, stack.Min())
				require.Equal(t, op.result, stack.ToArray())
			case stacksnqueues.POP:
				stack.Pop()
				require.Equal(t, op.min, stack.Min())
				require.Equal(t, op.result, stack.ToArray())
			}
		}
	}
}

func TestStackOfPlatesPush(t *testing.T) {
	testCases := []struct {
		stack stacksnqueues.StackOfPlates
		ops   []struct {
			typ              stacksnqueues.StackOpType
			value            int
			numberOfSubStack int
			result           []int
		}
	}{
		{
			stack: stacksnqueues.NewStackOfPlates(1),
			ops: []struct {
				typ              stacksnqueues.StackOpType
				value            int
				numberOfSubStack int
				result           []int
			}{
				{typ: stacksnqueues.PUSH, value: 1, numberOfSubStack: 1, result: []int{1}},
				{typ: stacksnqueues.PUSH, value: 2, numberOfSubStack: 2, result: []int{1, 2}},
				{typ: stacksnqueues.POP, numberOfSubStack: 1, result: []int{1}},
			},
		},
		{
			stack: stacksnqueues.NewStackOfPlates(2),
			ops: []struct {
				typ              stacksnqueues.StackOpType
				value            int
				numberOfSubStack int
				result           []int
			}{
				{typ: stacksnqueues.PUSH, value: 1, numberOfSubStack: 1, result: []int{1}},
				{typ: stacksnqueues.PUSH, value: 2, numberOfSubStack: 1, result: []int{1, 2}},
				{typ: stacksnqueues.POP, numberOfSubStack: 1, result: []int{1}},
				{typ: stacksnqueues.PUSH, value: -1, numberOfSubStack: 1, result: []int{1, -1}},
				{typ: stacksnqueues.PUSH, value: 0, numberOfSubStack: 2, result: []int{1, -1, 0}},
				{typ: stacksnqueues.POP, numberOfSubStack: 1, result: []int{1, -1}},
			},
		},
	}

	for _, testCase := range testCases {
		stack := testCase.stack
		for _, op := range testCase.ops {
			switch op.typ {
			case stacksnqueues.PUSH:
				stack.Push(op.value)
				require.Equal(t, op.numberOfSubStack, stack.NumberOfSubStack())
				require.Equal(t, op.result, stack.ToArray())
			case stacksnqueues.POP:
				stack.Pop()
				require.Equal(t, op.numberOfSubStack, stack.NumberOfSubStack())
				require.Equal(t, op.result, stack.ToArray())
			}
		}
	}
}

func TestMyQueue(t *testing.T) {
	testCases := []struct {
		queue stacksnqueues.MyQueue
		ops   []struct {
			typ    stacksnqueues.QueueOpType
			value  int
			result []int
		}
	}{
		{
			queue: stacksnqueues.MyQueue{},
			ops: []struct {
				typ    stacksnqueues.QueueOpType
				value  int
				result []int
			}{
				{typ: stacksnqueues.ENQUEUE, value: 1, result: []int{1}},
				{typ: stacksnqueues.ENQUEUE, value: 2, result: []int{1, 2}},
				{typ: stacksnqueues.DEQUEUE, value: 1, result: []int{2}},
				{typ: stacksnqueues.ENQUEUE, value: 1, result: []int{2, 1}},
				{typ: stacksnqueues.DEQUEUE, value: 2, result: []int{1}},
				{typ: stacksnqueues.DEQUEUE, value: 1, result: []int{}},
			},
		},
	}

	for _, testCase := range testCases {

		queue := testCase.queue
		for _, op := range testCase.ops {
			switch op.typ {
			case stacksnqueues.ENQUEUE:
				queue.Enqueue(op.value)
				require.Equal(t, op.result, queue.ToArray())
			case stacksnqueues.DEQUEUE:
				value := queue.Dequeue()
				require.Equal(t, op.value, value)
				require.Equal(t, op.result, queue.ToArray())
			}
		}
	}
}
