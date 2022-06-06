package prob

func DeleteCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			index1 := fast
			index2 := head
			for {
				if index1 == index2 {
					break
				}
				index1 = index1.Next
				index2 = index2.Next

			}
			return index2

		}
	}
	return nil

}
