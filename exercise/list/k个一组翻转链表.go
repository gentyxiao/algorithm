package list

// k个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Next: head,
	}
	end := dummy
	pre := dummy
	for end.Next != nil {
		// 一组最末端
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reversList(start)
		start.Next = next
		pre = start
		end = pre
	}
	return dummy.Next
}

// func reversList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	curr := head
// 	var pre *ListNode
// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = pre
// 		pre = curr
// 		curr = next
// 	}
// 	return pre
// }
