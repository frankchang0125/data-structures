package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Circular queue, which wastes an item of array
type Queue struct {
	front    int
	back     int
	capacity int
	Items    []int
}

func Constructor(capacity int) *Queue {
	if capacity <= 0 {
		return nil
	}

	return &Queue{
		capacity: capacity,
		Items:    make([]int, capacity),
	}
}

// Push pushes node to the back of queue,
// if queue is full, nothing will be pushed
func (q *Queue) Push(val int) {
	if q.IsFull() {
		return
	}

	q.back = (q.back + 1) % q.capacity
	q.Items[q.back] = val
}

// Pop pops node from front of queue,
// if queue is empty, -1 is returned
func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}

	item := q.Items[q.front]
	q.front = (q.front + 1) % q.capacity
	return item
}

func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Items[(q.front+1)%q.capacity]
}

func (q *Queue) Back() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Items[q.back]
}

func (q *Queue) IsFull() bool {
	return (q.back+1)%q.capacity == q.front
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.back
}

func (q *Queue) GetSize() int {
	size := q.back - q.front
	if size < 0 {
		size += q.capacity
	}
	return size
}

func (q *Queue) PrintQueue() {
	values := []string{}

	if !q.IsEmpty() {
		idx := q.front

		for {
			idx = (idx + 1) % q.capacity
			values = append(values, strconv.Itoa(q.Items[idx]))

			if idx == q.back {
				break
			}
		}
	}

	fmt.Printf("[%s]\n", strings.Join(values, ", "))
}

// Reference: https://goo.gl/dZ119Q
func main() {
	queue := Constructor(4)
	if queue == nil {
		return
	}

	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
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
	queue.PrintQueue()
	queue.Push(4)
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.Push(5)
	fmt.Println(queue.Front())
	fmt.Println(queue.Back())
	queue.Push(6)
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
