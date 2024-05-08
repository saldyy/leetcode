package main

import (
	"github.com/saldyy/leetcode/libs"
)

func Solution24(head *libs.Node) *libs.Node {
	current := head
	for current != nil && current.Next != nil {
		next := current.Next
		tmp := current.Value
		current.Value = next.Value
		next.Value = tmp
		current = next.Next
	}

	return head
}
