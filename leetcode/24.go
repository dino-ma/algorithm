package main

import (
	"fmt"
)

type ListNode struct{
	Next *ListNode
	Val int
}


func makeListNode(nums []int) *ListNode {
	if len(nums) == 0{
		return nil
	}

	res := &ListNode{
		Val:nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i],}
		temp = temp.Next
	}

	return  res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next == nil {
		return head
	}
	p := &ListNode{Next:head}
	h := p
	for p !=nil && p.Next != nil && p.Next.Next !=nil{
		a := p.Next
		b := a.Next
		a.Next = b.Next
		p.Next = b
		b.Next = a
		p = a
	}

	return h.Next
}

func main()  {

	a := makeListNode([]int{1,2,3,4})
	//for a != nil {
	//	fmt.Println(a.Val)
	//	a = a.Next
	//}
	result := swapPairs(a)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}