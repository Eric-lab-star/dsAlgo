package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNode   [2]*TreeNode
}

type KeyValue int

func (k KeyValue) LessThan(k1 KeyValue) bool { return k < k1 }

func InsertNode(tree **TreeNode, key KeyValue) {
	*tree, _ = InsertRNode(*tree, key)
}

func InsertRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return &TreeNode{KeyValue: key}, false
	}
	dir := 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNode[dir], done = InsertRNode(rootNode.LinkedNode[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}
	return BalanceTree(rootNode, dir), true
}

func BalanceTree(root *TreeNode, dir int) *TreeNode {
	node := root.LinkedNode[dir]
	checker := 2*dir - 1
	if node.BalanceValue == checker {
		node.BalanceValue = 0
		root.BalanceValue = 0
		return SingleRotation(root, dir)
	}
	adjustBalance(root, dir, checker)
	return doubleRotation(root, dir)
}

func SingleRotation(root *TreeNode, dir int) *TreeNode {
	saveNode := root.LinkedNode[dir]
	root.LinkedNode[dir] = saveNode.LinkedNode[1-dir]
	saveNode.LinkedNode[1-dir] = root
	return saveNode
}

func adjustBalance(rootNode *TreeNode, dir int, balanceValue int) {
	node := rootNode.LinkedNode[dir]
	oppNode := node.LinkedNode[1-dir]
	switch oppNode.BalanceValue {
	case 0: //RL, LR
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue: //RLL
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default: //LRR
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

func doubleRotation(rootNode *TreeNode, dir int) *TreeNode {
	//rotation --> rotation
	saveNode := rootNode.LinkedNode[dir].LinkedNode[1-dir]
	rootNode.LinkedNode[dir].LinkedNode[1-dir] = saveNode.LinkedNode[dir]
	saveNode.LinkedNode[dir] = rootNode.LinkedNode[dir]
	rootNode.LinkedNode[dir] = saveNode
	return SingleRotation(rootNode, dir)
}

func (tree *TreeNode) Print() {
	avlTree, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(avlTree))
}

func main() {
	var tree *TreeNode
	InsertNode(&tree, 10)
	InsertNode(&tree, 18)
	InsertNode(&tree, 15)
	tree.Print()
}
