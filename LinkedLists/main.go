package linkedlists

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Root *Node[T]
}

// Append inserts node at last index
func (l *LinkedList[T]) Append(next *Node[T]) {
	endNode := l.Root
	for endNode.Next != nil {
		endNode = endNode.Next
	}
	endNode.Next = next
}

func (l *LinkedList[T]) InsertAt(new *Node[T], index int) {
	if index <= 0 {
		new.Next = l.Root
		l.Root = new
		return
	}
	i := 0
	indexNode := l.Root
	for indexNode.Next != nil && i < index-1 {
		indexNode = indexNode.Next
		i++
	}
	if i == index-1 {
		new.Next = indexNode.Next
	}
	indexNode.Next = new
}

func (l *LinkedList[T]) DeleteAt(index int) {
	if index == 0 {
		l.Root = l.Root.Next
		return
	}
	i := 0
	indexNode := l.Root
	for indexNode.Next != nil && i < index-1 {
		indexNode = indexNode.Next
		i++
	}
	deleteNode := indexNode.Next
	if deleteNode == nil {
		return
	}
	nextNode := deleteNode.Next
	indexNode.Next = nextNode
}

func (l *LinkedList[T]) DeepCopy() *LinkedList[T] {
	return &LinkedList[T]{
		Root: deepCopy[T](l.Root),
	}
}

func (l *LinkedList[T]) Length() int {
	i := 0
	node := l.Root
	for node != nil {
		node = node.Next
		i++
	}
	return i
}

func deepCopy[T any](n *Node[T]) *Node[T] {
	if n == nil {
		return nil
	}
	newNode := &Node[T]{
		Data: n.Data,
		Next: deepCopy[T](n.Next),
	}
	return newNode
}

func (n *Node[T]) ToArray() []T {
	if n == nil {
		return make([]T, 0)
	}
	var ts []T
	node := n
	fmt.Println("Begining transfering from linked list to array")
	for node != nil {
		if node.Next != nil {
			fmt.Printf(">>: %v,\n", node.Data)
		} else {
			fmt.Printf(">>: %v.\n", node.Data)
		}
		ts = append(ts, node.Data)
		node = node.Next
	}
	return ts
}

// Remove Dups! Write code to remove duplicates from an unsorted linked list.
//
// # FOLLOW UP
//
// How would you solve this problem if a temporary buffer is not allowed?
func RemoveDups[T comparable](l *LinkedList[T]) *LinkedList[T] {
	result := &LinkedList[T]{
		Root: deepCopy[T](l.Root),
	}
	node := result.Root
	var preNode *Node[T]
	data := make(map[T]struct{})
	for node != nil {
		if _, exist := data[node.Data]; exist {
			// remove data
			preNode.Next = node.Next
		} else {
			data[node.Data] = struct{}{}
			preNode = node
		}
		node = node.Next
	}
	return result
}

// Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
func ReturnKthToLast[T any](l *LinkedList[T], k int) *LinkedList[T] {
	result := &LinkedList[T]{
		Root: deepCopy[T](l.Root),
	}
	node := result.Root
	i := 0
	for node != nil && i < k {
		node = node.Next
		i++
	}
	result.Root = node
	return result
}

// Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but
// the first and last node, not necessarily the exact middle) of a singly linked list, given only access to
// that node.
//
// # EXAMPLE
//
// lnput:the node c from the linked lista->b->c->d->e->f
//
// Result: nothing is returned, but the new linked list looks like a->b->d->e->f
func DeleteMiddleNode() {}

// Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
// before all nodes greater than or equal to x. If x is contained within the list, the values of x only need
// to be after the elements less than x (see below). The partition element x can appear anywhere in the
// "right partition"; it does not need to appear between the left and right partitions.
//
// # EXAMPLE
//
// Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition= 5]
//
// Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
func Partition() {}

// Sum Lists: You have two numbers represented by a linked list, where each node contains a single
// digit. The digits are stored in reverse order, such that the 1 's digit is at the head of the list. Write a
// function that adds the two numbers and returns the sum as a linked list.
//
// # EXAMPLE
//
// Input: (7-> 1 -> 6) + (5 -> 9 -> 2).That is, 617 + 295.
//
// Output: 2 -> 1 -> 9. That is, 912.
//
// # FOLLOW UP
//
// Suppose the digits are stored in forward order. Repeat the above problem.
//
// # EXAMPLE
//
// lnput: (6 -> 1 -> 7) + (2 -> 9 -> 5).That is, 617 + 295.
//
// Output: 9 -> 1 -> 2. That is, 912.
func SumLists() {}

// Palindrome: Implement a function to check if a linked list is a palindrome.
func Palindrome() {}

// Intersection: Given two (singly) linked lists, determine if the two lists intersect.
// Return the intersecting node. Note that the intersection is defined based on reference, not value.
// That is, if the kth node of the first linked list is the exact same node (by reference) as the jth
// node of the second linked list, then they are intersecting.
func Intersection() {}

// Loop Detection: Given a circular linked list, implement an algorithm that returns the node at the
// beginning of the loop.
//
// # DEFINITION
//
// Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so
// as to make a loop in the linked list.
//
// # EXAMPLE
//
// Input: A -> B -> C -> D -> E -> C [the same C as earlier]
//
// Output: C
func LoopDetection() {}
