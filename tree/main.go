package main

import (
	"fmt"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) Min() *Node {
	if tree.root == nil {
		return nil
	}

	for {
		if tree.root.left == nil {
			return tree.root
		}
		tree.root = tree.root.left
	}
}

func (tree *Tree) Max() *Node {
	if tree.root == nil {
		return nil
	}

	for {
		if tree.root.right == nil {
			return tree.root
		}
		tree.root = tree.root.right
	}

}

func (tree *Tree) Search(key int) bool {
	return search(tree.root, key)
}

func search(root *Node, key int) bool {
	if root == nil {
		return false
	}
	if root.key < key {
		return search(root.right, key)
	}
	if root.key > key {
		return search(root.left, key)
	}

	return true
}

func (tree *Tree) Insert(key int) {
	node := &Node{key, nil, nil}
	if tree.root == nil {
		tree.root = node
	} else {
		insert(tree.root, node)
	}
}

func insert(root, node *Node) {

	if node.key < root.key {
		if root.left == nil {
			root.left = node
		} else {
			insert(root.left, node)
		}
	}
	if node.key > root.key {
		if root.right == nil {
			root.right = node
		} else {
			insert(root.right, node)
		}
	}
}

func (tree *Tree) Remove(key int) {
	remove(tree.root, key)

}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		node.right = remove(node.right, key)
		return node
	}

	if node.key > key {
		node.left = remove(node.left, key)
		return node
	}

	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		return node.right
	}

	if node.right == nil {
		return node.left
	}

	leftmostright := node.right

	for {
		if leftmostright != nil && leftmostright.left != nil {
			leftmostright = leftmostright.left
		} else {
			break
		}
	}

	node.key = leftmostright.key
	node.right = remove(node.right, node.key)

	return node
}

func (tree *Tree) InOrder() {
	inorder(tree.root)
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Printf("%d \n", node.key)
		inorder(node.right)
	}
}

func (tree *Tree) PreOrder() {

	preorder(tree.root)
}

func preorder(node *Node) {
	if node != nil {
		fmt.Printf("%d\n", node.key)
		preorder(node.left)
		preorder(node.right)
	}

}

func (tree *Tree) String() {
	string(tree.root, 0)
}

func string(node *Node, level int) {
	msg := ""
	for i := 0; i < level; i++ {
		msg += "    "
	}
	msg += ">>"
	level += 1
	if node != nil {
		string(node.right, level)
		fmt.Printf(msg+"%d\n", node.key)
		string(node.left, level)
	}
}

func main() {
	tree := &Tree{}
	tree.Insert(8)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(12)
	tree.Insert(10)
	tree.String()
	tree.Remove(8)
	tree.String()
	check := tree.Search(8)
	fmt.Println(check)
	node := tree.Max()
	fmt.Println(node)

}
