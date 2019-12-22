package main

//第一次通过答案，分奇偶考虑，遍历时只需一个指针，o(n)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//second为待返回的节点，curr为指向当前节点的指针
	curr, second := head, head.Next

	//遍历依次，直到找到最后一组待交换的值（此时指针在该组的前一组值上）
	for curr.Next.Next != nil && curr.Next.Next.Next != nil {
		curr.Next.Next, curr.Next, curr = curr, curr.Next.Next.Next, curr.Next.Next
	}

	//当链表节点为偶数时
	if curr.Next.Next == nil {
		curr.Next.Next, curr.Next = curr, nil
	//当链表节点为奇数时
	} else{
		curr.Next.Next, curr.Next = curr, curr.Next.Next
	}
	return second

}
