package main

import "fmt"

type Node struct {
	prop int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddToHead(prop int) {
	node := &Node{prop: prop}
	if list.head != nil {
		node.next = list.head
		list.head.prev = node
	}
	list.head = node
}

func (list *LinkedList) NodeWith(prop int) *Node {
	for node := list.head; node != nil; node = node.next {
		if node.prop == prop {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddAfter(target, prop int) {
	node := &Node{prop: prop}
	nodeWith := list.NodeWith(target)

	if nodeWith != nil {
		node.next = nodeWith.next
		node.prev = nodeWith
		nodeWith.next.prev = node
		nodeWith.next = node
	}
}

func (list *LinkedList) LastNode() *Node {
	for node := list.head; node != nil; node = node.next {
		if node.next == nil {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddToEnd(prop int) {
	last := list.LastNode()
	node := &Node{prop: prop}

	if last != nil {
		last.next = node
		node.prev = last
	}
}

func (list *LinkedList) IterateToNext() {
	for node := list.head; node != nil; node = node.next {
		fmt.Println(node)
	}
}

func (list *LinkedList) NodeBetween(prev, next int) *Node {
	for node := list.head; node != nil; node = node.next {
		if node.next != nil && node.prev != nil {
			if node.next.prop == next && node.prev.prop == prev {
				return node
			}
		}
	}
	return nil
}
func main() {

	list := &LinkedList{}
	list.AddToHead(1)
	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(4)
	fmt.Println(list.NodeBetween(2, 4))
}
