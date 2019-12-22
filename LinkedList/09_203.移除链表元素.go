package main

//使用哑结点，o(n)
func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	curr, prec := head, dummy
	for curr != nil {
		if curr.Val == val{
			prec.Next = curr.Next
		} else{
			prec = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}
