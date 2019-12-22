package main

//代码简化
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next !=nil {
		fast, slow=fast.Next.Next, slow.Next
		if fast == slow{
			return true
		}
	}
	return false
}

//第一次代码，双指针
func hasCycle_(head *ListNode) bool {
	if head == nil{
		return false
	}
	fast, slow := head, head
	for  {
		if fast == nil || fast.Next == nil{return false}
		fast, slow=fast.Next.Next, slow.Next
		if fast == slow{
			return true
		}
	}
}

