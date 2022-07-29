package prob

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
