package list

// 环形链表 II，返回环的入口
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 1.确认快慢指针第一次相交的地方
	// 2.慢指针回到链表头部
	// 3.快慢指针改为每次移动一次，再次相遇时就是环的入口
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
