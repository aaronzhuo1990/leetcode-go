package datastruct

// ListNode defines for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeSingleLinkedList uses the given int slice to create a single linked list
// Please note that the head also contains value.
func MakeSingleLinkedList(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}
