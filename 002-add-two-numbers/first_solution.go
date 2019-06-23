package main

import (
	"fmt"
	"github.com/azhuox/leetcode-go/000datastruct"
)

func main() {
	l1 := makeListNode([]int{2, 4, 3})
	l2 := makeListNode([]int{6, 6, 6})

	result := addTwoNumbers(l1, l2)
	fmt.Printf("result:") // Should be 8 -> 0 -> 0 -> 1
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")
}

func makeListNode(is []int) *datastruct.ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &datastruct.ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &datastruct.ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func addTwoNumbers(l1 *datastruct.ListNode, l2 *datastruct.ListNode) *datastruct.ListNode {
	head := &datastruct.ListNode{} // Head's Val doesn't save data
	cur := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &datastruct.ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return head.Next
}
