package libs

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func AddNode(head *Node, value int) *Node {
	node := &Node{Value: value}
	if head == nil {
		return node
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	return head
}

func (n *Node) Add(val int) *Node {
	node := &Node{Value: val}
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	return n
}

func (n *Node) Print() {
	current := n
	for current != nil {
		fmt.Printf("%d->", current.Value)
		current = current.Next
	}
	fmt.Printf("\n")
}
