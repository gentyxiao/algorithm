package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转链表
func reversList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
