package stacksnqueues

type Stack []int

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	l := len(*s)
	value := (*s)[l-1]
	*s = (*s)[:l-1]
	return value
}

func (s *Stack) Peek() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

type StackOpType int

const (
	PUSH StackOpType = 1
	POP  StackOpType = 2
	PEEK StackOpType = 3
)

type StackType int

const (
	FIRST  StackType = 1
	SECOND StackType = 2
	THRID  StackType = 3
)

type QueueOpType int

const (
	ENQUEUE QueueOpType = 1
	DEQUEUE QueueOpType = 2
)

// Three in One: Describe how you could use a single array to implement three stacks.
type ThreeStack struct {
	stack          []int
	lenStackFirst  int
	lenStackSecond int
	lenStackThird  int
}

func NewThreeStack() ThreeStack {
	return ThreeStack{
		stack: make([]int, 0),
	}
}

func NewThreeStackFull(stack []int, lenFirst, lenSecond, lenThird int) ThreeStack {
	return ThreeStack{
		stack:          stack,
		lenStackFirst:  lenFirst,
		lenStackSecond: lenSecond,
		lenStackThird:  lenThird,
	}
}

func (t *ThreeStack) Pop(typ StackType) int {
	var value int
	var index int
	switch typ {
	case FIRST:
		index = t.lenStackFirst - 1
		t.lenStackFirst--
	case SECOND:
		index = t.lenStackFirst + t.lenStackSecond - 1
		t.lenStackSecond--
	case THRID:
		index = t.lenStackFirst + t.lenStackSecond + t.lenStackThird - 1
		t.lenStackThird--
	}

	temp := make([]int, index)
	copy(temp, t.stack[:index])
	value = t.stack[index]
	temp = append(temp, t.stack[index+1:]...)
	t.stack = temp
	return value
}

func (t *ThreeStack) Push(typ StackType, value int) {
	var index int
	switch typ {
	case FIRST:
		index = t.lenStackFirst
		t.lenStackFirst++
	case SECOND:
		index = t.lenStackFirst + t.lenStackSecond
		t.lenStackSecond++
	case THRID:
		index = t.lenStackFirst + t.lenStackSecond + t.lenStackThird
		t.lenStackThird++
	}
	temp := make([]int, index)
	copy(temp, t.stack[:index])
	temp = append(temp, value)
	temp = append(temp, t.stack[index:]...)
	t.stack = temp
}

func (t *ThreeStack) Peek(typ StackType) int {
	var value int
	switch typ {
	case FIRST:
		value = t.stack[t.lenStackFirst-1]
	case SECOND:
		value = t.stack[t.lenStackFirst+t.lenStackSecond-1]
	case THRID:
		value = t.stack[t.lenStackFirst+t.lenStackSecond+t.lenStackThird-1]
	}
	return value
}

func (t *ThreeStack) IsEmpty(typ StackType) bool {
	var isEmpty bool
	switch typ {
	case FIRST:
		isEmpty = t.lenStackFirst == 0
	case SECOND:
		isEmpty = t.lenStackSecond == 0
	case THRID:
		isEmpty = t.lenStackThird == 0
	}
	return isEmpty
}

func (t *ThreeStack) ToArray() []int {
	return t.stack
}

// Stack Min: How would you design a stack which, in addition to push and pop, has a function min
// which returns the minimum element? Push, pop and min should all operate in 0(1) time.
type StackMin struct {
	stack         Stack
	stackMinState Stack
	len           int
	min           int
}

func (s *StackMin) Push(value int) {
	s.stack.Push(value)
	if s.stackMinState.Length() > 0 && value < s.stackMinState.Peek() {
		s.stackMinState.Push(value)
	} else if s.stackMinState.Length() == 0 {
		s.stackMinState.Push(value)
	} else {
		s.stackMinState.Push(s.stackMinState.Peek())
	}
	s.len++
}

func (s *StackMin) Pop() int {
	s.stackMinState.Pop()
	s.len--
	return s.stack.Pop()
}

func (s *StackMin) Min() int {
	return s.stackMinState.Peek()
}

func (t *StackMin) ToArray() []int {
	return t.stack
}

// Stack of Plates: Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
// Therefore, in real life, we would likely start a new stack when the previous stack exceeds some
// threshold. Implement a data structure SetOfStacks that mimics this. SetOfStacks should be
// composed of several stacks and should create a new stack once the previous one exceeds capacity.
// SetOfStacks. push() and SetOfStacks. pop() should behave identically to a single stack
// (that is, pop () should return the same values as it would if there were just a single stack).
//
// # FOLLOW UP
//
// Implement a function popAt (int index) which performs a pop operation on a specific sub-stack.
type StackOfPlates struct {
	stacks    []Stack
	threshold int
}

func NewStackOfPlates(threshold int) StackOfPlates {
	if threshold < 1 {
		threshold = 1
	}
	return StackOfPlates{
		threshold: threshold,
	}
}

func (s *StackOfPlates) Push(value int) {
	if len(s.stacks) == 0 {
		s.stacks = append(s.stacks, Stack{})
	}
	lastIndex := len(s.stacks) - 1
	lenLastStack := len(s.stacks[lastIndex])

	if lenLastStack == s.threshold {
		newStack := Stack{}
		newStack.Push(value)
		s.stacks = append(s.stacks, newStack)
	} else {
		s.stacks[lastIndex].Push(value)
	}
}

func (s *StackOfPlates) Pop() int {
	l := len(s.stacks)
	lastIndex := len(s.stacks) - 1
	value := s.stacks[lastIndex].Pop()
	lenLastStack := len(s.stacks[lastIndex])

	if lenLastStack == 0 && l > 1 {
		s.stacks = s.stacks[:l-1]
	}

	return value
}

func (s *StackOfPlates) NumberOfSubStack() int {
	return len(s.stacks)
}

func (s *StackOfPlates) ToArray() []int {
	var result []int
	for _, stack := range s.stacks {
		result = append(result, stack...)
	}
	return result
}

// Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks.
type MyQueue struct {
	stack  Stack
	stack1 Stack
}

func (q *MyQueue) Enqueue(value int) {
	q.stack.Push(value)
}

func (q *MyQueue) Dequeue() int {
	for q.stack.Length() > 0 {
		temp := q.stack.Pop()
		q.stack1.Push(temp)
	}
	value := q.stack1.Pop()

	for q.stack1.Length() > 0 {
		temp := q.stack1.Pop()
		q.stack.Push(temp)
	}

	return value
}

func (q *MyQueue) ToArray() []int {
	return q.stack
}

// Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use
// an additional temporary stack, but you may not copy the elements into any other data structure
// (such as an array). The stack supports the following operations: push, pop, peek, and is Empty.

// Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly "first in, first
// out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
// or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
// that type). They cannot select which specific animal they would like. Create the data structures to
// maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
// and dequeueCat. You may use the built-in Linked list data structure.
