package main

import "fmt"

type Node struct {
	key   int
	right *Node
	left  *Node
}

type Tree struct {
	root *Node
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
	if root.key < node.key {
		if root.right == nil {
			root.right = node
		} else {
			insert(root.right, node)
		}
	} else {
		if root.left == nil {
			root.left = node
		} else {
			insert(root.left, node)
		}
	}
}

func (tree *Tree) String() {
	fmt.Println("========")
	string(tree.root, 0)
	fmt.Println("========")
}

func string(node *Node, level int) {
	msg := ""
	for i := 0; i < level; i++ {
		msg += "    "
	}
	msg += ">>"
	level++
	if node != nil {
		string(node.right, level)
		fmt.Printf(msg+"%d\n", node.key)
		string(node.left, level)
	}
}

func (tree *Tree) Search(key int) bool {
	if tree.root != nil {
		return search(tree.root, key)
	}

	return false
}

func search(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if node.key < key {
		return search(node.right, key)
	}

	if node.key > key {
		return search(node.left, key)
	}

	return true

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

	if node.right == nil {
		return node.left
	}
	if node.left == nil {
		return node.right
	}

	if node.left == nil && node.right == nil {
		return nil
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

func (tree *Tree) Min() (node *Node) {
	if tree.root == nil {
		return nil
	}
	root := tree.root
	for {
		if root.left != nil {
			root = root.left
		} else {
			break
		}
	}

	return root

}

func (tree *Tree) Max() (node *Node) {
	if tree.root == nil {
		return nil
	}
	root := tree.root
	for {
		if root.right != nil {
			root = root.right
		} else {
			break
		}
	}

	return root
}

func main() {
	tree := &Tree{}
	tree.Insert(8)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(10)
	tree.String()
	max := tree.Max()
	fmt.Println(max)
	tree.String()
}

