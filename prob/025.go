package prob

import "github.com/emirpasic/gods/stacks/arraystack"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := arraystack.New(), arraystack.New()
	for l1 != nil {
		s1.Push(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2.Push(l2.Val)
		l2 = l2.Next
	}

	carry, dummy := 0, new(ListNode)
	for !s1.empty() || !s2.empty() || carry > 0 {
		v1, ok := s1.Pop()
		if ok {
			carry += v1.(int)
		}
		v2, ok := s2.Pop()
		if ok {
			carry += v2.(int)
		}
		node := &ListNode{Val: carry % 10, Next: dummy.Next}
		dummy.Next = node
		carry /= 10
	}
	return dummy.Next
}
