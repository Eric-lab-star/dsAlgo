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

	fmt.Println("=====")
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

func Remove(tree **treeNode, key int) {
	*tree, _ = removeRNode(*tree, key)
}

func removeRNode(root *treeNode, key int) (*treeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root.Key == key {
		switch {
		case root.LinkedNodes[0] == nil:
			return root.LinkedNodes[1], false
		case root.LinkedNodes[1] == nil:
			return root.LinkedNodes[0], false
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
	// return nil, false
}

func removeBalance(rootNode *treeNode, dir int) (*treeNode, bool) {
	node := rootNode.LinkedNodes[1-dir]
	checker := 2*dir - 1
	switch node.BalanceValue {
	case checker:
		balancer(rootNode, checker, dir)
		return doubleRotation(rootNode, 1-dir), false
	case -checker:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, 1-dir), false
	}
	rootNode.BalanceValue = -checker
	node.BalanceValue = checker
	return singleRotation(rootNode, dir), true
}

func main() {
	var tree *treeNode
	Insert(&tree, 40)
	Insert(&tree, 45)
	Insert(&tree, 50)
	Insert(&tree, 60)
	tree.Print()
	Remove(&tree, 40)
	tree.Print()
}
