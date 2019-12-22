package main

//题解，将当前节点值设为下一节点值，删除下一节点
func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}