package prob

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromed(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
