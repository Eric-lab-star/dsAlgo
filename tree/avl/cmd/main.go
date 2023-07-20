package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type treeNode struct {
	Key          int
	BalanceValue int
	LinkedNodes  [2]*treeNode
}

func (tree *treeNode) Print() {
	avl, err := json.MarshalIndent(tree, " ", "  ")
	if err != nil {
		log.Panicf("Marshal Json error %v", err)
		return
	}
	fmt.Println(string(avl))
}

func Insert(tree **treeNode, key int) {
	*tree, _ = insertRNode(*tree, key)
}

func insertRNode(tree *treeNode, key int) (*treeNode, bool) {
	if tree == nil {
		return &treeNode{Key: key}, false
	}
	dir := 0
	if tree.Key < key {
		dir = 1
	}
	var done bool
	tree.LinkedNodes[dir], done = insertRNode(tree.LinkedNodes[dir], key)
	if done {
		return tree, true
	}
	tree.BalanceValue = tree.BalanceValue + (2*dir - 1)
	switch tree.BalanceValue {
	case 0:
		return tree, true
	case 1, -1:
		return tree, false
	}
	return balanceTree(tree, dir), true
}

func balanceTree(root *treeNode, dir int) *treeNode {
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

func singleRotation(root *treeNode, dir int) *treeNode {
	save := root.LinkedNodes[dir]
	root.LinkedNodes[dir] = save.LinkedNodes[1-dir]
	save.LinkedNodes[1-dir] = root
	return save
}

func balancer(root *treeNode, checker int, dir int) {
	node := root.LinkedNodes[dir]
	save := node.LinkedNodes[1-dir]
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

func doubleRotation(root *treeNode, dir int) *treeNode {
	node := root.LinkedNodes[dir]
	save := node.LinkedNodes[1-dir]
	node.LinkedNodes[1-dir] = save.LinkedNodes[dir]
	save.LinkedNodes[dir] = node
	root.LinkedNodes[dir] = save
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
