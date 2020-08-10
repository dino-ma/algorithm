package main

import "fmt"

type ListNode struct{
	Next *ListNode
	Val int
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	for head != nil && head.Next != nil && fast !=nil && fast.Next != nil  {
		if head  == fast{
			return true
		}
		head = head.Next
		fast = fast.Next.Next
	}


	return false
}


func main()  {


	//make cycle
	head := &ListNode{Val:0}
	next := &ListNode{Val:1}
	head.Next = next
	second := &ListNode{Val:2}
	next.Next = second
	second.Next = head

	fmt.Println(hasCycle(head))


}
