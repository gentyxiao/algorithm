package list

// k个一组升序链表合并
// 顺序合并，拆分为将两个链表合并为一个链表
// 依次将每个链表合并
func merge(listA, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	if listB == nil {
		return listA
	}
	newList := &ListNode{}
	head := newList
	for listA != nil && listB != nil {
		if listA.Val < listB.Val {
			newList.Next = listA
			listA = listA.Next
		} else {
			newList.Next = listB
			listB = listB.Next
		}
		newList = newList.Next
	}
	if listA != nil {
		newList.Next = listA
	} else {
		newList.Next = listB
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	var list *ListNode
	for i := 0; i < len(lists); i++ {
		list = merge(list, lists[i])
	}
	return list
}
