package list

// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 示例 2：

// 输入：l1 = [], l2 = []
// 输出：[]
// 示例 3：

// 输入：l1 = [], l2 = [0]
// 输出：[0]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}
	tmp := dummy
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			tmp.Next = list2
			list2 = list2.Next
			tmp = tmp.Next
		} else {
			tmp.Next = list1
			list1 = list1.Next
			tmp = tmp.Next
		}
	}
	if list1 != nil {
		tmp.Next = list1
	}

	if list2 != nil {
		tmp.Next = list2
	}
	return dummy.Next
}
