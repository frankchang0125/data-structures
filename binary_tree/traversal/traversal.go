package main

import (
	"fmt"
)

type BinaryTree struct {
	root *Node
}

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	val    int
}

// Pre-order traversal: Middle -> Left -> Right: DFS, use recursive (stack) to implement
func PreOrder(node *Node) {
	if node != nil {
		fmt.Println(node.val)
		PreOrder(node.left)
		PreOrder(node.right)
	}
}

// In-order traversal: Left -> Middle -> Right
// For Binary Search tree, we can sort the numbers
// from small to large by in-order traversal
func InOrder(node *Node) {
	if node != nil {
		InOrder(node.left)
		fmt.Println(node.val)
		InOrder(node.right)
	}
}

// leftMost find the left-most node of node's subtree
func leftMost(node *Node) *Node {
	if node == nil {
		return nil
	}

	current := node

	for {
		if current.left == nil {
			return current
		}
		current = current.left
	}
}

// For node: X's in-order successor:
//  If node: X has right child node:
//  	=> Its in-order successor must be the left-most node
//         of its right child node
//  If node: X doesn't have right child node:
//		a. If node: X is the left node of its parent node:
//			=> Its in-order successor is its parent node
//      b. If node: X is not the left node of its parent node:
//			=> Its in-order successor is the first node: Y's parent,
//      	   which node: Y is the first-found left child node traverse
//			   along node: X's ancestors
//
//      	ex:         a
//      	        /       \
//      	       b         c
//      	    /     \   /     \
//      	   d       e f       g
//      	          / \
//      	         h   i
//
//      	The successor node of node: i in inorder traversal is node: a.
//      	(We can find a parent node: b, which is the left node of
//      	 its parent node: a)
func InOrderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	current := node

	if current.right != nil {
		return leftMost(current.right)
	}

	ancestor := current.parent

	for {
		if ancestor == nil || ancestor.left == current {
			break
		}
		current = ancestor
		ancestor = ancestor.parent
	}
	return ancestor
}

// rightMost find the right-most node of node's subtree
func rightMost(node *Node) *Node {
	if node == nil {
		return nil
	}

	current := node

	for {
		if current.right == nil {
			return current
		}
		current = current.right
	}
}

// For node: X's in-order predecessor:
//  If node: X has left child node:
//      Its in-order predecessor must be the right-most node
//      of its left child node
//  If node: X doesn't have left child node:
//      Its in-order predecessor must be the first node: Y's parent
//      which node: Y is the first-found right child node traverse along
//      node: X's ancestors
func InOrderPredecessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	current := node

	if current.left != nil {
		return rightMost(current.left)
	}

	ancestor := current.parent

	for {
		if ancestor == nil || ancestor.right == current {
			break
		}
		current = ancestor
		ancestor = ancestor.parent
	}
	return ancestor
}

// Post-order traversal: Left -> Right -> Middle
func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.left)
		PostOrder(node.right)
		fmt.Println(node.val)
	}
}

// Level-order traversal: BFS, use queue to implement
func LevelOrder(node *Node) {
	queue := make([]*Node, 0)

	if node != nil {
		queue = append(queue, node)
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:] // Dequeue
		fmt.Println(currentNode.val)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left) // Enqueue
		}

		if currentNode.right != nil {
			queue = append(queue, currentNode.right) // Enqueue
		}
	}
}

// References:
//  https://goo.gl/CQ3wN4
//  https://goo.gl/XYndzM
func main() {
	rootNode := &Node{val: 50}
	nodeA := &Node{val: 20}
	nodeB := &Node{val: 60}
	nodeC := &Node{val: 10}
	nodeD := &Node{val: 40}
	nodeE := &Node{val: 30}
	nodeF := &Node{val: 70}
	nodeG := &Node{val: 80}

	rootNode.left = nodeA
	nodeA.parent = rootNode
	rootNode.right = nodeB
	nodeB.parent = rootNode

	nodeA.left = nodeC
	nodeC.parent = nodeA
	nodeA.right = nodeD
	nodeD.parent = nodeA

	nodeB.right = nodeF
	nodeF.parent = nodeB

	nodeD.left = nodeE
	nodeE.parent = nodeD

	nodeF.right = nodeG
	nodeG.parent = nodeF

	tree := &BinaryTree{
		root: rootNode,
	}

	fmt.Println("Pre-order traversal:")
	PreOrder(tree.root)
	fmt.Println()

	fmt.Println("In-order traversal:")
	InOrder(tree.root)
	fmt.Println()

	fmt.Println("Post-order traversal:")
	PostOrder(tree.root)
	fmt.Println()

	fmt.Println("Level-order traversal:")
	LevelOrder(tree.root)
	fmt.Println()

	var predecessor, successor *Node

	predecessor = InOrderPredecessor(rootNode)
	if predecessor == nil {
		fmt.Printf("%d's predecessor not found\n", rootNode.val)
	} else {
		fmt.Printf("%d's predecessor: %d\n", rootNode.val, predecessor.val)
	}

	successor = InOrderSuccessor(rootNode)
	if successor == nil {
		fmt.Printf("%d's successor not found\n", rootNode.val)
	} else {
		fmt.Printf("%d's successor: %d\n", rootNode.val, successor.val)
	}
	fmt.Println()

	predecessor = InOrderPredecessor(nodeC)
	if predecessor == nil {
		fmt.Printf("%d's predecessor not found\n", nodeC.val)
	} else {
		fmt.Printf("%d's predecessor: %d\n", nodeC.val, predecessor.val)
	}

	successor = InOrderSuccessor(nodeC)
	if successor == nil {
		fmt.Printf("%d's successor not found\n", nodeC.val)
	} else {
		fmt.Printf("%d's successor: %d\n", nodeC.val, successor.val)
	}
	fmt.Println()

	predecessor = InOrderPredecessor(nodeE)
	if predecessor == nil {
		fmt.Printf("%d's predecessor not found\n", nodeE.val)
	} else {
		fmt.Printf("%d's predecessor: %d\n", nodeE.val, predecessor.val)
	}

	successor = InOrderSuccessor(nodeE)
	if successor == nil {
		fmt.Printf("%d's successor not found\n", nodeE.val)
	} else {
		fmt.Printf("%d's successor: %d\n", nodeE.val, successor.val)
	}
	fmt.Println()

	predecessor = InOrderPredecessor(nodeG)
	if predecessor == nil {
		fmt.Printf("%d's predecessor not found\n", nodeG.val)
	} else {
		fmt.Printf("%d's predecessor: %d\n", nodeG.val, predecessor.val)
	}

	successor = InOrderSuccessor(nodeG)
	if successor == nil {
		fmt.Printf("%d's successor not found\n", nodeG.val)
	} else {
		fmt.Printf("%d's successor: %d\n", nodeG.val, successor.val)
	}
}
