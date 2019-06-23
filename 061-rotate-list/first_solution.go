package main

import (
	"fmt"
	"github.com/azhuox/leetcode-go/000datastruct"
)

func main() {

	input := datastruct.MakeSingleLinkedList([]int{1, 2, 3, 4, 5})
	result := rotateRight(input, 2)

	fmt.Printf("result: ") // Should be 4 -> 5 -> 1 -> 2 -> 3
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")
}

func rotateRight(head *datastruct.ListNode, k int) *datastruct.ListNode {
	if k == 0 || head == nil {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail.Next == nil {
			return rotateRight(head, k%(i+1)) // i + 1 is the length of the list
		}
		tail = tail.Next
	}

	newTail := head
	for tail.Next != nil {
		newTail, tail = newTail.Next, tail.Next
	}

	tail.Next = head
	newhead := newTail.Next
	newTail.Next = nil

	return newhead
}
