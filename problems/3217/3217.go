package problem3217

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMap(nums []int) map[int]bool {
	m := make(map[int]bool)
	fmt.Println("Nums:", nums)
	for _, v := range nums {
		m[v] = true
	}

	return m
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := getMap(nums)

	for head != nil {
		if !m[head.Val] {
			break
		}
		head = head.Next
	}

	cur := head

	for cur.Next != nil {
		if m[cur.Next.Val] {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}

	return head
}

func Execute(a []int, head *ListNode) *ListNode {

	return modifiedList(a, head)
}
