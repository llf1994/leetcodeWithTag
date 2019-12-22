package main

type ListNode struct {
	Val  int
	Next *ListNode
}
//递归
func reverseList(head *ListNode) *ListNode {
	//首先找到尾节点（通过递归）
	if head == nil || head.Next == nil{ //head == nil 检查空链表
		return head
	}
	//当尾节点返回时，head为尾节点的前一个节点
	tail := reverseList(head.Next)

	//让尾节点指向前一个节点head，head指向nil，此时head为新的尾节点
	head.Next.Next, head.Next = head, nil

	//在第一次递归中，head为第一个节点;在第n次递归中，head为第n个节点
	//整个递归流程不对tail操作，tail不断上抛，直到第一次递归，并抛出
	return tail
}


//遍历
func reverseList_(head *ListNode) *ListNode {
	//单链表需要保存前一个节点
	var prec *ListNode //初始化一个指针，指针值为nil；var prec ListNode和prec:=new(ListNode)都会将ListNode的初值设为零值；
	next, curr := new(ListNode), head
	for curr != nil {
		//先取得下一个节点
		next = curr.Next
		//将当前节点指向前一个节点
		curr.Next = prec
		//保存当前节点
		prec = curr
		//跳到下一个节点
		curr = next
	}
	return prec
}
