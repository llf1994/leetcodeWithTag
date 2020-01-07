package main

//第一次通过答案，o(n)
func deleteDuplicates_II(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	curr, prec := head, dummy

	for curr != nil && curr.Next != nil{
		//如果出现相同值
		if curr.Val == curr.Next.Val{
			//找到第一个非重复值所在节点
			tmp := curr.Val
			curr = curr.Next.Next
			for curr != nil && curr.Val == tmp{
				curr = curr.Next
			}
			prec.Next = curr
		} else{
			curr, prec = curr.Next, curr
		}
	}
	return dummy.Next
}