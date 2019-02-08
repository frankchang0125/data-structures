package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	prev *Node
	next *Node
	val  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.getNode(index)
	if node == nil {
		return -1
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		val: val,
	}

	node := this.head
	if this.head == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.next = this.head
		this.head = newNode
		node.prev = this.head
	}

	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		val: val,
	}

	if this.tail == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.prev = this.tail
		this.tail.next = newNode
		this.tail = newNode
	}

	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	} else if index == this.size {
		this.AddAtTail(val)
		return
	}

	node := this.getNode(index)
	newNode := &Node{
		prev: node.prev,
		next: node,
		val:  val,
	}

	node.prev.next = newNode
	node.prev = newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.size-1 {
		return
	}

	if index == 0 {
		this.head = this.head.next
		this.head.prev = nil
		this.size--
		return
	}

	if index == this.size-1 {
		this.tail = this.tail.prev
		this.tail.next = nil
		this.size--
		return
	}

	node := this.getNode(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	this.size--
}

/** Reverse the entire linked list **/
func (this *MyLinkedList) Reverse() {
	curNode := this.head
	this.head = this.tail
	this.tail = curNode

	for curNode != nil {
		nextNode := curNode.next
		curNode.next = curNode.prev
		curNode.prev = nextNode
		curNode = nextNode
	}
}

/** Print the entire linked list from head to end **/
func (this *MyLinkedList) PrintList() {
	node := this.head
	values := []string{}

	for node != nil {
		values = append(values, strconv.Itoa(node.val))
		node = node.next
	}

	fmt.Printf("[%s]\n", strings.Join(values, ", "))
}

func (this *MyLinkedList) getNode(index int) *Node {
	if index > this.size-1 {
		return nil
	}

	var node *Node

	if index <= this.size/2 {
		// Iterate from head
		node = this.head
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		// Iterate from tail
		node = this.tail
		for i := this.size - 1; i > index; i-- {
			node = node.prev
		}
	}

	return node
}

func main() {
	list := Constructor()
	list.AddAtHead(10)
	list.AddAtHead(20)
	list.AddAtIndex(4, 30)
	list.AddAtIndex(2, 40)
	list.DeleteAtIndex(1)
	list.DeleteAtIndex(4)
	fmt.Println(list.Get(5))
	fmt.Println(list.Get(1))
	list.AddAtTail(50)
	list.AddAtIndex(0, 60)
	list.AddAtIndex(4, 70)
	list.PrintList()
	list.Reverse()
	list.PrintList()
}
