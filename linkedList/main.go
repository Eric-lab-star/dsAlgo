package main

import "fmt"

type Node struct {
	prop     int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (list *LinkedList) AddToHead(prop int) {
	node := &Node{prop: prop}
	if list.headNode != nil {
		node.nextNode = list.headNode
	}
	list.headNode = node
}

func (list *LinkedList) NodeWithValue(prop int) *Node {
	for node := list.headNode; node != nil; node = node.nextNode {
		if node.prop == prop {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddAfter(target, prop int) {
	node := &Node{prop: prop}
	for nodeWith := list.headNode; nodeWith != nil; nodeWith = nodeWith.nextNode {
		if nodeWith.prop == target {
			node.nextNode = nodeWith.nextNode
			nodeWith.nextNode = node
		}
	}

}

func (list *LinkedList) LastNode() *Node {
	for node := list.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddToLast(prop int) {
	lastNode := list.LastNode()
	node := &Node{prop: prop}
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (list *LinkedList) Iterate() {
	for node := list.headNode; node != nil; node = node.nextNode {
		fmt.Println(node)
	}
}
func main() {
	list := &LinkedList{}
	list.AddToHead(5)
	list.AddAfter(5, 2)
	list.AddToLast(3)
	list.Iterate()

}
