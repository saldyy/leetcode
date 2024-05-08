//https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/
package main

import (
	"math/big"
	"strconv"

	"github.com/saldyy/leetcode/libs"
)

func Solution2816(head *libs.Node) *libs.Node {
	num := big.NewInt(0)
	current := head
	for current != nil {
		num = num.Mul(num, big.NewInt(10))
		num = num.Add(num, big.NewInt(int64(current.Value)))
		current = current.Next
	}

	strNum := big.NewInt(0).Mul(num, big.NewInt(2)).String()
	first, _ := strconv.Atoi(string(strNum[0]))
	newNode := &libs.Node{Value: first}
	current = newNode
	for i := 1; i < len(strNum); i++ {
		n, _ := strconv.Atoi(string(strNum[i]))
		newNode.Next = &libs.Node{Value: n}
		newNode = newNode.Next
	}
	return current
}
