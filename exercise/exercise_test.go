package exercise

import (
	"algorithm/exercise/list"
	"fmt"
	"testing"
)

func TestThreeSum1(t *testing.T) {
	t.Log("aaaaaaaaaaaaa")
	node1 := &list.ListNode{
		Val: 1,
	}
	node2 := &list.ListNode{
		Val: 8,
	}
	node3 := &list.ListNode{
		Val: 3,
	}
	node4 := &list.ListNode{
		Val: 6,
	}
	node5 := &list.ListNode{
		Val: 5,
	}
	node6 := &list.ListNode{
		Val: 4,
	}
	node7 := &list.ListNode{
		Val: 7,
	}
	node8 := &list.ListNode{
		Val: 2,
	}
	node9 := &list.ListNode{
		Val: 9,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node7.Next = node8
	node8.Next = node9

	list.SortList(node1)
	for node1 != nil {
		fmt.Println(node1.Val)
		node1 = node1.Next
	}
}
