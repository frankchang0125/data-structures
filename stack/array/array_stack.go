package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	Items []int
}

func Constructor() *Stack {
	return &Stack{
		Items: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.Items = append(s.Items, val)
}

// Pop returns popped value, if stack is empty, return -1
func (s *Stack) Pop() int {
	if s.GetSize() == 0 {
		return -1
	}

	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return s.GetSize() == 0
}

// Top returns top value in stack, if stack is empty, return -1
func (s *Stack) Top() int {
	if s.GetSize() == 0 {
		return -1
	}

	return s.Items[len(s.Items)-1]
}

func (s *Stack) GetSize() int {
	return len(s.Items)
}

func (s *Stack) PrintStack() {
	values := []string{}

	for i := len(s.Items) - 1; i >= 0; i-- {
		values = append(values, strconv.Itoa(s.Items[i]))
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
