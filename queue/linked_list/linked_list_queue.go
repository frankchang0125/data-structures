package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Queue struct {
	front *Node
	back  *Node
	size  int
}

type Node struct {
	next *Node
	val  int
}

func Constructor() *Queue {
	return &Queue{}
}

// Push pushes node to the back of queue
func (q *Queue) Push(val int) {
	newNode := &Node{
		val: val,
	}

	if q.IsEmpty() {
		q.back = newNode
		q.front = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}

	q.size++
}

// Pop pops node from front of queue,
// if queue is empty, -1 is returned
func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}

	node := q.front
	q.front = node.next
	node.next = nil

	if q.front == nil {
		q.back = nil
	}

	q.size--
	return node.val
}

func (q *Queue) Front() int {
	if q.front == nil {
		return -1
	}
	return q.front.val
}

func (q *Queue) Back() int {
	if q.back == nil {
		return -1
	}
	return q.back.val
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) GetSize() int {
	return q.size
}

func (q *Queue) PrintQueue() {
	values := []string{}
	node := q.front

	for node != nil {
		values = append(values, strconv.Itoa(node.val))
		node = node.next
	}

	fmt.Printf("[%s]\n", strings.Join(values, ", "))
}

func main() {
	queue := Constructor()
	queue.Push(1)
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.Push(2)
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.Push(3)
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.PrintQueue()
	queue.Pop()
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.Pop()
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.Pop()
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.PrintQueue()
}
