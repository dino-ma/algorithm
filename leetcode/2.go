package main



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode = nil
	for head != nil {
		head.Next,pre,head = pre,head,head.Next
	}

	return pre
}

func main()  {

}
