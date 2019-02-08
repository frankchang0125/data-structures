package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	head *Node
	size int
}

type Node struct {
	next *Node
	val  int
}

func Constructor() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.head = &Node{
		next: s.head,
		val:  val,
	}
	s.size++
}

// Pop returns popped value, if stack is empty, return -1
func (s *Stack) Pop() int {
	if s.head == nil {
		return -1
	}

	node := s.head
	s.head = s.head.next
	s.size--
	return node.val
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

// Top returns top value in stack, if stack is empty, return -1
func (s *Stack) Top() int {
	if s.head == nil {
		return -1
	}
	return s.head.val
}

func (s *Stack) GetSize() int {
	return s.size
}

func (s *Stack) PrintStack() {
	values := []string{}
	node := s.head

	for node != nil {
		values = append(values, strconv.Itoa(node.val))
		node = node.next
	}

	fmt.Printf("[%s]\n", strings.Join(values, ", "))
}

func main() {
	stack := Constructor()
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	stack.PrintStack()
	stack.Push(2)
	fmt.Println(stack.Top())
	stack.Push(3)
	fmt.Println(stack.GetSize())
	stack.PrintStack()
	fmt.Println(stack.Pop())
	stack.PrintStack()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetSize())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	stack.PrintStack()
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.GetSize())
	stack.PrintStack()
}
