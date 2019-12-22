package main

//......题目给的是排序链表，直接遍历即可，o(n)
func deleteDuplicates(head *ListNode) *ListNode {
	curr, pre := head,head
	for curr != nil{
		curr=curr.Next
		if curr == nil {break}

		if curr.Val == pre.Val{
			pre.Next = curr.Next
		} else{
			pre = curr
		}
	}
	return head
}

//哈希表，o(n)
func deleteDuplicates_(head *ListNode) *ListNode {
	tmp, curr, pre := make(map[int]bool), head, new(ListNode)

	for curr != nil{
		if _, ok := tmp[curr.Val]; ok{
			pre.Next = pre.Next.Next
		} else{
			tmp[curr.Val] = true
			//只有元素不重复时才更新pre
			pre = curr
		}
		//第一个元素不存在重复
		curr =  curr.Next
	}
	return head
}