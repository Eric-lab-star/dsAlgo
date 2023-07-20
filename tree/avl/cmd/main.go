package main

import (
	"encoding/json"
	"fmt"
)

type treeNode struct {
	Key          int
	BalanceValue int
	LinkedNode   [2]*treeNode
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func (tree *treeNode) Print() {
	avl, err := json.MarshalIndent(tree, "  ", "    ")
	errHandler(err)
	fmt.Println(string(avl))
}

func Insert(tree **treeNode, key int) {
	*tree, _ = insertRNode(*tree, key)
}

func insertRNode(root *treeNode, key int) (*treeNode, bool) {
	if root == nil {
		return &treeNode{Key: key}, false
	}
	dir := 0
	if root.Key < key {
		dir = 1
	}
	var done bool
	root.LinkedNode[dir], done = insertRNode(root.LinkedNode[dir], key)
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
	return balanceTree(root, dir), false
}

func balanceTree(root *treeNode, dir int) *treeNode {
	checker := 2*dir - 1
	node := root.LinkedNode[dir]
	if node.BalanceValue == checker {
		node.BalanceValue = 0
		root.BalanceValue = 0
		return singleRotation(root, dir)
	}
	balancer(root, checker, dir)
	return doubleRotation(root, dir)
}

func balancer(root *treeNode, checker int, dir int) {
	save := root.LinkedNode[dir].LinkedNode[1-dir]
	node := root.LinkedNode[dir]
	switch save.BalanceValue {
	case 0:
		node.BalanceValue = 0
		root.BalanceValue = 0
	case checker:
		node.BalanceValue = 0
		root.BalanceValue = -checker
	default:
		node.BalanceValue = checker
		root.BalanceValue = 0
	}
	save.BalanceValue = 0
}

func singleRotation(root *treeNode, dir int) *treeNode {
	save := root.LinkedNode[dir]
	root.LinkedNode[dir] = save.LinkedNode[1-dir]
	save.LinkedNode[1-dir] = root
	return save
}

func doubleRotation(root *treeNode, dir int) *treeNode {
	save := root.LinkedNode[dir].LinkedNode[1-dir]
	root.LinkedNode[dir].LinkedNode[1-dir] = save.LinkedNode[dir]
	save.LinkedNode[dir] = root.LinkedNode[dir]
	root.LinkedNode[dir] = save
	return singleRotation(root, dir)
}

func main() {
	var tree *treeNode
	Insert(&tree, 45)
	Insert(&tree, 40)
	Insert(&tree, 80)
	Insert(&tree, 50)
	Insert(&tree, 90)
	Insert(&tree, 60)
	tree.Print()

}
