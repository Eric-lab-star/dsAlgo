package main

import (
	"encoding/json"
	"fmt"
)

type treeNode struct {
	Key          int
	BalanceValue int
	LinkedNodes  [2]*treeNode
}

func InsertNode(tree **treeNode, key int) {
	*tree, _ = insertRNode(*tree, key)
}

func insertRNode(node *treeNode, key int) (*treeNode, bool) {
	if node == nil {
		node = &treeNode{Key: key}
		return node, false
	}
	dir := 0
	if node.Key < key {
		dir = 1
	}
	var done bool
	node.LinkedNodes[dir], done = insertRNode(node.LinkedNodes[dir], key)
	if done {
		return node, true
	}
	node.BalanceValue = node.BalanceValue + (2*dir - 1)
	switch node.BalanceValue {
	case 0:
		return node, false
	case 1, -1:
		return node, false
	}
	return BalanceTree(node, dir), true
}

func BalanceTree(root *treeNode, dir int) *treeNode {
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

func singleRotation(root *treeNode, checker int) *treeNode {
	save := root.LinkedNodes[checker]
	root.LinkedNodes[checker] = save.LinkedNodes[1-checker]
	save.LinkedNodes[1-checker] = root
	return save
}

func balancer(root *treeNode, checker int, dir int) {
	node := root.LinkedNodes[dir]
	switch node.LinkedNodes[1-dir].BalanceValue {
	case 0:
		node.BalanceValue = 0
		root.BalanceValue = 0
	case checker:
		node.BalanceValue = -checker
		root.BalanceValue = 0
	default:
		node.BalanceValue = 0
		root.BalanceValue = checker
	}
	node.LinkedNodes[1-dir].BalanceValue = 0
}

func doubleRotation(root *treeNode, dir int) *treeNode {
	save := root.LinkedNodes[dir].LinkedNodes[1-dir]
	root.LinkedNodes[dir].LinkedNodes[1-dir] = save.LinkedNodes[dir]
	save.LinkedNodes[dir] = root.LinkedNodes[dir]
	root.LinkedNodes[dir] = save
	save = singleRotation(root, dir)
	return save
}

func (tree *treeNode) Print() {
	avl, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(avl))
}

func main() {
	var tree *treeNode
	InsertNode(&tree, 10)
	InsertNode(&tree, 9)
	InsertNode(&tree, 2)
	tree.Print()

}
