package prob

func isPalindrome(head *ListNode) bool {
	mid := MiddleNode(head)
	tmp := mid.Next
	mid.Next = nil
	tmp = ReverseList(tmp)
	for head != nil && tmp != nil {
		if head.Val != tmp.Val {
			return false
		}
		head = head.Next
		tmp = tmp.Next

	}
	return true
}

func MiddleNodePre(head *ListNode) *ListNode {
	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre
}
