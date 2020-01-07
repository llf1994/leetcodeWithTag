package main

//第一次通过答案，单指针加计数器，o(n)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//用计数器count记录链表节点，curr停留在尾节点的前一个节点
	curr, second, count := head, head.Next, 0
	for curr.Next.Next != nil {
		curr.Next, curr = curr.Next.Next, curr.Next
		count++
	}

	//如果链表节点为奇数(curr移动的次数count为节点数-2)
	if count%2 != 0 {
		curr.Next.Next, curr.Next = second, nil
	//如果链表节点为偶数
	} else {
		curr.Next, curr.Next.Next = second, nil
	}

	return head
}
