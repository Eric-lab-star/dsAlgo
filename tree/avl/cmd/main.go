package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

type KeyValue int

func (key KeyValue) LessThan(k KeyValue) bool { return key < KeyValue(k) }

func (key KeyValue) EqualTo(k KeyValue) bool { return key == KeyValue(k) }

func (tree *TreeNode) Print() {
	avlTree, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(avlTree))
}

func (tree *TreeNode) InsertNode(key KeyValue) {
	*tree = TreeNode{KeyValue: key}
}

func InsertNode(tree **TreeNode, key KeyValue) {
	*tree, _ = insertRNode(*tree, key)
}

func insertRNode(root *TreeNode, key KeyValue) (*TreeNode, bool) {
	if root == nil {
		root = &TreeNode{KeyValue: key}
		return root, false
	}
	return nil, false
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{}
	tree.InsertNode(10)
	tree.Print()
}

