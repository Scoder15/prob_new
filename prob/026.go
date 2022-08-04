package prob

func reorderList(head *ListNode) {
	mid := MiddleNode(head)
	tmp := mid.Next
	mid.Next = nil
	tmp = ReverseList(tmp)
	head = MergeTwoLists(head, tmp)
}

func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for l1 != nil && l2 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy
}
