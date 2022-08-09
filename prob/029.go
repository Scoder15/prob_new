package prob

func insert(aNode *ListNode, x int) *ListNode {
	node := new(ListNode)
	node.Val = x
	node.Next = nil
	if aNode == nil {
		aNode = node
		return aNode
	}
	p := aNode
	for {
		if (p.Val <= x && x <= p.Next.Val) || ((p.Val > p.Next.Val) && (x >= p.Val) || (x <= p.Next.Val)) || p.Next == aNode {
			node.Next = p.Next
			p.Next = node
			break
		}
		p = p.Next

	}
	return aNode
}
