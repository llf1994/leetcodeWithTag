package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node:=dummy
	for l1 != nil && l2 != nil{
		if l1.Val > l2.Val{
			node.Next,l2 = l2,l2.Next
			node=node.Next
		} else {
			node.Next, l1 = l1,l1.Next
			node=node.Next
		}
	}

	if l1 == nil{
		node.Next = l2
	} else {
		node.Next = l1
	}

	return dummy.Next
}