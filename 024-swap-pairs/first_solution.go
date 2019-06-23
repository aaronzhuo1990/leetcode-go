package main

import (
	"fmt"
	"github.com/azhuox/leetcode-go/000datastruct"
)

func main() {
	input := datastruct.MakeSingleLinkedList([]int{1, 2, 3, 4})
	result := swapPairs(input)

	fmt.Printf("result:") // Should be 2 -> 1 -> 4 -> 3
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")
}

func swapPairs(head *datastruct.ListNode) *datastruct.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Take {1, 2, 3, 4,} as an example, think about the case where head points 3
	// new head points to 4. swapPairs(newHead.Next) returns nil.
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
