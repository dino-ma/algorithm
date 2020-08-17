package main



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil  {
		return nil
	}
	p := &ListNode{Next:head}
	h := p

	for p != nil&& p.Next !=nil &&p.Next.Next !=nil {
		a := p.Next
		b := a.Next
		p.Next = b
		a.Next = b.Next
		b.Next = a
		p = a
	}

	return h.Next
}

func main(){

}
