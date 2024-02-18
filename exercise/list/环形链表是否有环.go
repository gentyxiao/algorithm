package list

// 环形链表，检查链表是否有环
func hasCycle(head *ListNode) bool {
	// 快慢指针实现
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
}
