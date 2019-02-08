package main

import (
	"fmt"
)

type MinStack struct {
	Data *Stack
	Min  *Stack
}

func Constructor() *MinStack {
	return &MinStack{
		Data: StackConstructor(),
		Min:  StackConstructor(),
	}
}

func (s *MinStack) Push(val int) {
	if s.Min.IsEmpty() {
		s.Min.Push(val)
	} else {
		// Only push min stack if insert value is smaller or equal
		// to the top value of min stack
		min := s.Min.Top()
		if val <= min {
			s.Min.Push(val)
		}
	}

	s.Data.Push(val)
}

func (s *MinStack) Pop() int {
	top := s.Data.Pop()

	// Only pop min stack if top value of data stack
	// is equal to the top value of min stack
	if s.Min.Top() == top {
		s.Min.Pop()
	}
	return top
}

func (s *MinStack) IsEmpty() bool {
	return s.Data.IsEmpty()
}

func (s *MinStack) Top() int {
	return s.Data.Top()
}

func (s *MinStack) GetSize() int {
	return s.Data.GetSize()
}

func (s *MinStack) GetMin() int {
	return s.Min.Top()
}

func (s *MinStack) PrintStack() {
	s.Data.PrintStack()
}

func (s *MinStack) PrintMinStack() {
	s.Min.PrintStack()
}

// References:
//  1. https://goo.gl/ghQX2Q
//  2. https://goo.gl/6u8vMN
func main() {
	stack := Constructor()
	stack.Push(5)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Push(3)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Push(4)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Push(8)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Push(1)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Push(10)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Push(9)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}
