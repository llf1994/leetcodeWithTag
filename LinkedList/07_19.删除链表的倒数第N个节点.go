package main

//题解，使用哑结点避开临界值

//双指针一次遍历,o(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//如果链表只有一个节点
	if head.Next == nil{
		var a *ListNode
		return a
	}

	slow, fast, count := head, head, 0
	//循环结束后slow为被删除节点的前一个节点
	for fast != nil{
		fast = fast.Next
		count++
		if count - n > 1{
			slow = slow.Next
		}
	}
	//如果删除的是第一个节点
	if count-n == 0 {return head.Next}

	slow.Next = slow.Next.Next
	return head
}

//第一次答案，两次遍历，o(n)
func removeNthFromEnd_(head *ListNode, n int) *ListNode {
	//如果链表只有一个节点
	if head.Next == nil{
		var a *ListNode
		return a
	}
	curr, count  := head, 0

	//统计链表长度
	for curr != nil {
		count++
		curr = curr.Next
	}
	curr = head

	//如果删除的是第一个节点
	if count - n == 0 {return head.Next}

	//删除倒数第n个节点,拿到倒数第n-1个节点
	for i:=1;i< count-n ;i++{
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return head
}