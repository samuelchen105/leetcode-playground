package problem

// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// d 1 2 3 4 5
// | |

// d 1 2 3 4 5
// |     |

// d 1 2 3 4 5
//   |     |

// d 1 2 3 4 5
//     |     |

// d 1 2 3 4 5 n
//       |     |

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	anchor := head

	for n > 0 {
		anchor = anchor.Next
		n--
	}

	for anchor != nil {
		anchor = anchor.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next

	return dummy.Next
}
