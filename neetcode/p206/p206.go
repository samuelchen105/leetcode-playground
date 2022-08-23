package problem

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}
