package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Set struct {
	items []int
}

func Constructor(capacity int) *Set {
	if capacity == 0 {
		return nil
	}

	items := make([]int, capacity)
	for i := range items {
		items[i] = -1
	}

	return &Set{
		items: items,
	}
}

func (s *Set) FindSet(x int) int {
	if s.items[x] < 0 {
		// Root found, return
		return x
	}

	// Recursive calls to find the root
	root := s.FindSet(s.items[x])
	// Set collapsing (path compression)
	s.items[x] = root
	return root
}

func (s *Set) UnionSet(x, y int) {
	rootX, rootY := s.FindSet(x), s.FindSet(y)

	if s.items[rootX] <= s.items[rootY] {
		// x's set has equal or more elements than y's set,
		// merge y's set into x's set and update y's set root
		// to x's set root
		s.items[rootX] += s.items[rootY]
		s.items[rootY] = rootX
	} else {
		// y's set has more elements than x's set,
		// merge x's set into y's set and update x's set root
		// to y's set root
		s.items[rootY] += s.items[rootX]
		s.items[rootX] = rootY
	}
}

func (s *Set) PrintSet() {
	values := make([]string, len(s.items))

	for i, item := range s.items {
		values[i] = strconv.Itoa(item)
	}

	fmt.Printf("[%s]\n", strings.Join(values, ", "))
}

// References:
//  1. https://goo.gl/9xHh8f
//  2. https://goo.gl/4EP2us
func main() {
	set := Constructor(6)
	set.PrintSet()
	set.UnionSet(1, 2)
	set.PrintSet()
	set.UnionSet(0, 2)
	set.PrintSet()
	set.UnionSet(3, 5)
	set.PrintSet()
	set.UnionSet(2, 5)
	set.PrintSet()
}
