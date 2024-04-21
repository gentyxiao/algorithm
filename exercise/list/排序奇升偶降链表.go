package list

// 给定一个奇数位升序，偶数位降序的链表，将其重新排序，要求时间O(n)空间O(1)的排序
// 输入: 1->8->3->6->5->4->7->2->NULL 输出: 1->2->3->4->5->6->7->8->NULL
//
func SortList(head *ListNode) *ListNode {
	// 第一步：拆分出基、偶链表
	tmp := head
	var l1, l2 *ListNode = &ListNode{}, &ListNode{}
	l1Tmp := l1
	l2Tmp := l2
	for i := 1; tmp != nil; i++ {
		if i%2 == 1 {
			// 基数链表
			l1Tmp.Next = tmp
			l1Tmp = l1Tmp.Next
		} else {
			// 偶数链表
			l2Tmp.Next = tmp
			l2Tmp = l2Tmp.Next
		}
		tmp = tmp.Next
	}
	// 第二步：翻转偶数链表
	l3 := reversList(l2.Next)
	// var prev *ListNode
	// var curr *ListNode = l2
	// for curr != nil {
	// 	next := curr.Next
	// 	curr.Next = prev
	// 	prev = curr
	// 	curr = next
	// }
	// // 偶数链表翻转后的链表
	// l3 := prev

	// 第三步：合并两个有序链表，l1 与 l3，得到l4
	l4 := mergeTwoLists(l1.Next, l3)
	// var l4 *ListNode = &ListNode{}
	// tmp = l4
	// for l1 != nil && l3 != nil {
	// 	if l1.Val >= l3.Val {
	// 		tmp.Next = l3
	// 		l3 = l3.Next
	// 		tmp = tmp.Next
	// 	} else {
	// 		tmp.Next = l1
	// 		l1 = l1.Next
	// 		tmp = tmp.Next
	// 	}
	// }
	// // 连接剩下的链表节点
	// if l1 != nil {
	// 	tmp.Next = l1
	// }
	// if l3 != nil {
	// 	tmp.Next = l3
	// }
	return l4
}
