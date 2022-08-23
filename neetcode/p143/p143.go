package problem

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// find middle
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse second half
	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for second != nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	// merge
	first, second := head, prev
	for first != nil && second != nil {
		tmpF, tmpS := first.Next, second.Next

		second.Next = first.Next
		first.Next = second

		first = tmpF
		second = tmpS
	}
}
