package main

import "fmt"

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	val    int
}

// In-order traversal: Left -> Middle -> Right
// For Binary search tree, we can sort the numbers
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
//      Its in-order successor must be the left-most node
//      of its right child node
//  If node: X doesn't have right child node:
//      Its in-order successor must be the first node: Y's parent
//      which node: Y is the first-found left child node traverse along
//      node: X's ancestors
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

func (t *BinarySearchTree) Insert(val int) {
	node := &Node{val: val}

	if t.root == nil {
		t.root = node
		return
	}

	current := t.root

	for {
		if val <= current.val {
			if current.left == nil {
				node.parent = current
				current.left = node
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				node.parent = current
				current.right = node
				return
			}
			current = current.right
		}
	}
}

// For node with two child nodes:
//  Its in-order predecessor or successor must be a:
//      a. Leaf node
//      b. A node with only one child node
//  Its because for node with two child nodes:
//      Its in-order predecessor must be the:
//          Largest-value node of the left subtree
//      Its in-order successor must be the:
//          Smallest-value node of the right subtree
func (t *BinarySearchTree) Delete(val int) {
	found, node := t.Search(val)
	if !found {
		// Node to be deleted is not in the tree, do nothing
		return
	}

	parent := node.parent

	if node.left == nil && node.right == nil {
		// Node to be deleted has no child node
		if parent != nil {
			if node == parent.left {
				// Node to be deleted is a left node of its parent
				parent.left = nil
			} else {
				// Node to be deleted is a right node of its parent
				parent.right = nil
			}
		} else {
			// Node to be deleted is a root node
			t.root = nil
		}
	} else if node.left != nil && node.right == nil {
		// Node to be deleted has only one left child node
		if parent != nil {
			if node == parent.left {
				// Node to be deleted is a left node of its parent
				parent.left = node.left
			} else {
				// Node to be deleted is a right node of its parent
				parent.right = node.left
			}
		} else {
			// Node to be deleted is a root node
			t.root = node.left
		}
	} else if node.left == nil && node.right != nil {
		// Node to be deleted has only one right child node
		if parent != nil {
			if node == parent.left {
				// Node to be deleted is a left node of its parent
				parent.left = node.right
			} else {
				// Node to be deleted is a right node of its parent
				parent.right = node.right
			}
		} else {
			// Node to be deleted is a root node
			t.root = node.right
		}
	} else {
		// Node to be deleted has two child nodes,
		// pick either in-order predecessor child node
		// or in-order successor child node to replace it
		predecessor := InOrderPredecessor(node)

		// Predecessor must have only left child node
		// Otherwise, its right child node would be the predecessor
		// of the node to be deleted instead
		predecessorChild := predecessor.left
		if predecessorChild != nil {
			predecessorChild.parent = predecessor.parent
		}

		predecessorParent := predecessor.parent
		if predecessor == predecessorParent.left {
			predecessorParent.left = predecessorChild
		} else {
			predecessorParent.right = predecessorChild
		}

		// Don't actually delete the node,
		// but replace its value with predecessor's value instead
		node.val = predecessor.val
	}
}

func (t *BinarySearchTree) Search(target int) (found bool, node *Node) {
	current := t.root

	for {
		if current == nil {
			return false, nil
		} else if current.val == target {
			return true, current
		}

		if target < current.val {
			current = current.left
		} else if target > current.val {
			current = current.right
		}
	}
}

// For Binary search tree, we can sort the numbers
// from small to large by in-order traversal
func (t *BinarySearchTree) Sort() {
	InOrder(t.root)
}

// References:
//  https://goo.gl/bEnBTi
//  https://goo.gl/gHNsnf
//  https://goo.gl/CQ3wN4
func main() {
	t := &BinarySearchTree{}
	t.Insert(50)
	t.Insert(20)
	t.Insert(30)
	t.Insert(40)
	t.Insert(10)
	t.Insert(80)
	t.Insert(70)
	t.Insert(90)
	t.Insert(30)
	t.Insert(100)
	t.Insert(60)
	LevelOrder(t.root)
	fmt.Println()
	t.Sort()
	fmt.Println()

	target := 60
	found, _ := t.Search(target)
	if !found {
		fmt.Printf("%d not found\n", target)
	} else {
		fmt.Printf("%d found\n", target)
	}

	target = 30
	found, _ = t.Search(target)
	if !found {
		fmt.Printf("%d not found\n", target)
	} else {
		fmt.Printf("%d found\n", target)
	}

	target = 70
	found, _ = t.Search(target)
	if !found {
		fmt.Printf("%d not found\n", target)
	} else {
		fmt.Printf("%d found\n", target)
	}

	target = 110
	found, _ = t.Search(target)
	if !found {
		fmt.Printf("%d not found\n", target)
	} else {
		fmt.Printf("%d found\n", target)
	}

	target = 0
	found, _ = t.Search(target)
	if !found {
		fmt.Printf("%d not found\n", target)
	} else {
		fmt.Printf("%d found\n", target)
	}
	fmt.Println()

	_, ts := t.Search(70)
	a := InOrderPredecessor(ts)
	if a == nil {
		fmt.Println("70 has no predecessor")
	} else {
		fmt.Printf("70's predecessor: %d\n", a.val)
	}
	b := InOrderSuccessor(ts)
	if b == nil {
		fmt.Println("70 has no successor")
	} else {
		fmt.Printf("70's successor: %d\n", b.val)
	}
	fmt.Println()

	t.Delete(50)
	LevelOrder(t.root)
	fmt.Println()
	t.Delete(20)
	LevelOrder(t.root)
	fmt.Println()
	t.Delete(80)
	LevelOrder(t.root)
	fmt.Println()
}
