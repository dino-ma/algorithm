package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversLinkList(head *ListNode) *ListNode {
	var pre *ListNode = nil

	for head != nil {
		next := head.Next   // 先存下下一个节点，不然一会就没了
		head.Next = pre     // 当前节点指向上一个节点
		pre = head
		head = next
	}

	return pre
}

func showLinkList(head *ListNode)   {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	head := &ListNode{Val:0}
	tail := head
	for i :=1; i<10 ;i++  {
		tail.Next = & ListNode{Val:i}
		tail = tail.Next
	}


	showLinkList(head)
	fmt.Println("resvers")
	resversList := reversLinkList(head)
	showLinkList(resversList)
}

