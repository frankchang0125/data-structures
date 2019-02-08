package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MyLinkedList struct {
	head *Node
}

type Node struct {
	next *Node
	val  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.head

	for i := 0; i < index; i++ {
		if node == nil {
			return -1
		}
		node = node.next
	}

	if node == nil {
		return -1
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{
		next: this.head,
		val:  val,
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &Node{
			next: nil,
			val:  val,
		}
		return
	}

	node := this.head

	for {
		if node.next == nil {
			node.next = &Node{
				next: nil,
				val:  val,
			}
			return
		}
		node = node.next
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.head == nil {
		if index == 0 {
			this.head = &Node{
				next: nil,
				val:  val,
			}
			return
		}
		return
	}

	node := this.head
	var i int

	// Iterate to (index - 1) node, if possible
	for i = 0; i < index-1; i++ {
		if node.next != nil {
			node = node.next
		} else {
			return
		}
	}

	node.next = &Node{
		next: node.next,
		val:  val,
	}

	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.head == nil {
		return
	}

	node := this.head
	var i int

	// Iterate to (index - 1) node, if possible
	for i = 0; i < index-1; i++ {
		if node.next != nil {
			node = node.next
		} else {
			return
		}
	}

	if node.next == nil {
		return
	}

	node.next = node.next.next
}

/** Reverse the entire linked list **/
func (this *MyLinkedList) Reverse() {
	prevNode := this.head
	var curNode, nextNode *Node

	if this.head == nil {
		// Empty linked list
		return
	}

	if this.head.next == nil {
		// Linked list with only one node
		return
	}

	prevNode = nil
	curNode = this.head
	nextNode = this.head.next

	for {
		if nextNode == nil {
			// curNode is now pointing to the last node,
			// reverse current node and update head pointer
			curNode.next = prevNode
			this.head = curNode
			return
		}

		// Reverse current node
		curNode.next = prevNode

		// Move each pointer to its next node
		prevNode = curNode
		curNode = nextNode
		nextNode = nextNode.next
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
