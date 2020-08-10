package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	//如果head为空或者head的指针为空则当前链表无环
	if head == nil || head.Next == nil {
		return nil
	}
	//初始化快慢指针，
	slow ,fast := head, head
	//以快指针为主（减少循环次数）
	for fast != nil &&fast.Next != nil && fast.Next.Next !=nil {
		//如果快指针没有后继 则当前链表不存在环
		if fast.Next ==nil {
			return nil
		}
		slow = slow.Next//每次走一步
		fast = fast.Next.Next//每次走两步
		//如果快慢指针相遇
		if slow == fast {
			//寻找相交的头节点，并且将快指针改变为每次只走一步
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}

	}

	return nil
}


func main(){
	//make cycle ii
	head := &ListNode{Val:3}
	next := &ListNode{Val:2}
	head.Next = next
	second := &ListNode{Val:0}
	next.Next = second

	third := &ListNode{Val:4}
	second.Next = third
	third.Next = next


	result := detectCycle(head)
	fmt.Println(result.Val)

}
