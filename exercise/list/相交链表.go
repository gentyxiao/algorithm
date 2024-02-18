package list

// 相交链表，找出相交链链表的起始节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap := map[*ListNode]*ListNode{}
	for headA != nil {
		hashMap[headA] = headA
		headA = headA.Next
	}
	for headB != nil {
		val, ok := hashMap[headB]
		if ok {
			return val
		}
		headB = headB.Next
	}
	return nil
}
