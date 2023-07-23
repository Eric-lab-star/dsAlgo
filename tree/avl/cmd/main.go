package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Key          int
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

func (t *TreeNode) Print() {
	avl, err := json.MarshalIndent(t, " ", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(avl))
}

func Insert(tree **TreeNode, key int) {
	*tree, _ = insertRNode(*tree, key)
}

func insertRNode(root *TreeNode, key int) (*TreeNode, bool) {
	if root == nil {
		return &TreeNode{Key: key}, false
	}
	dir := 0
	if root.Key < key {
		dir = 1
	}
	var done bool
	root.LinkedNodes[dir], done = insertRNode(root.LinkedNodes[dir], key)
	if done {
		return root, true
	}
	root.BalanceValue = root.BalanceValue + (2*dir - 1)
	switch root.BalanceValue {
	case 0:
		return root, true
	case 1, -1:
		return root, false
	}
	return balanceTree(root, dir), true
}

func balanceTree(root *TreeNode, dir int) *TreeNode {
	node := root.LinkedNodes[dir]
	checker := 2*dir - 1
	if node.BalanceValue == checker {
		node.BalanceValue = 0
		root.BalanceValue = 0
		return singleRotation(root, dir)
	}
	balancer(root, checker, dir)
	return doubleRotation(root, dir)
}

func balancer(root *TreeNode, checker int, dir int) {
	node := root.LinkedNodes[dir]
	save := node.LinkedNodes[1-dir]
	switch save.BalanceValue {
	case 0:
		node.BalanceValue = 0
		root.BalanceValue = 0
	case checker:
		node.BalanceValue = 0
		root.BalanceValue = -checker
	case -checker:
		node.BalanceValue = checker
		root.BalanceValue = 0
	}
	save.BalanceValue = 0
}

func doubleRotation(root *TreeNode, dir int) *TreeNode {
	node := root.LinkedNodes[dir]
	save := node.LinkedNodes[1-dir]
	node.LinkedNodes[1-dir] = save.LinkedNodes[dir]
	save.LinkedNodes[dir] = node
	root.LinkedNodes[dir] = save
	return singleRotation(root, dir)
}

func singleRotation(root *TreeNode, dir int) *TreeNode {
	save := root.LinkedNodes[dir]
	root.LinkedNodes[dir] = save.LinkedNodes[1-dir]
	save.LinkedNodes[1-dir] = root
	return save
}

func Remove(tree **TreeNode, key int) {
	*tree, _ = removeRNode(*tree, key)
}

func removeRNode(root *TreeNode, key int) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}

	if root.Key == key {
		switch {
		case root.LinkedNodes[1] == nil:
			return root.LinkedNodes[0], false
		case root.LinkedNodes[0] == nil:
			return root.LinkedNodes[1], false
		}
		heir := root.LinkedNodes[0]
		for heir.LinkedNodes[1] != nil {
			heir = heir.LinkedNodes[1]
		}
		root.Key = heir.Key
		key = heir.Key
	}
	dir := 0
	if root.Key < key {
		dir = 1
	}
	var done bool
	root.LinkedNodes[dir], done = removeRNode(root.LinkedNodes[dir], key)
	if done {
		return root, true
	}
	root.BalanceValue = root.BalanceValue + (1 - 2*dir)
	switch root.BalanceValue {
	case 0:
		return root, false
	case 1, -1:
		return root, true
	}
	return removeBalance(root, dir)
}
func removeBalance(root *TreeNode, dir int) (*TreeNode, bool) {
	fmt.Println("remove")
	node := root.LinkedNodes[1-dir]
	checker := 2*dir - 1
	switch node.BalanceValue {
	case -checker:
		node.BalanceValue = 0
		root.BalanceValue = 0
		return singleRotation(root, 1-dir), true
	case checker:
		balancer(root, checker, 1-dir)
		return doubleRotation(root, 1-dir), true
	}
	root.BalanceValue = -checker
	node.BalanceValue = checker
	return singleRotation(root, 1-dir), true

}

func main() {
	var tree *TreeNode

	Insert(&tree, 50)
	Insert(&tree, 40)
	Insert(&tree, 60)
	Insert(&tree, 55)
	Insert(&tree, 70)
	tree.Print()
	Remove(&tree, 40)
	tree.Print()

}
