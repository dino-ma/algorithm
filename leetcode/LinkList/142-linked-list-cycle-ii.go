package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow , fast := head ,head

	for fast !=nil && fast.Next !=nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for fast != head{
				fast = fast.Next
				head = head.Next
			}
			return head
		}
	}
	return nil

}

func main()  {

}