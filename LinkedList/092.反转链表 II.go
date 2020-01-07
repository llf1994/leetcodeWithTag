package main

//第一次通过答案，单指针遍历+3个记录点和1个计数器，o(n)
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	curr, pre, count := head, head, 1

	//p1、p3保存被反转部分节点的前一个节点和后一个节点，p2保存被反转部分中的最后一个节点
	p1, p2, p3 := new(ListNode), new(ListNode), new(ListNode)
	p1.Next = head
	for curr != nil{
		//被反转部分外的前一个节点
		if count == m-1 {p1=curr}
		//被反转部分中的最后一个节点
		if count == n {p2,p3=curr,curr.Next}

		//进入反转部分
		if count >= m && count <= n{
			curr, curr.Next, pre = curr.Next, pre, curr
			count++
			continue
		}
		curr = curr.Next
		count++

		if count > n {break}
	}

	p1.Next, p1.Next.Next = p2, p3

	//如果从head开始反转，返回p2
	if m == 1{return p2}
	return head
}