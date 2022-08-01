package prob

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
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
