package main

import (
	"fmt"
	"github.com/azhuox/leetcode-go/000datastruct"
)

func main() {

	input := datastruct.MakeSingleLinkedList([]int{1, 1, 2, 2, 3, 4, 4, 5})
	result := removeDuplicates(input)

	fmt.Printf("result: ") // Should be 3 -> 5
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")
}

func removeDuplicates(list *datastruct.ListNode) *datastruct.ListNode {

	if list == nil || list.Next == nil {
		return list
	}

	first := list
	for first.Val == first.Next.Val {
		first = first.Next
	}
	if first != list {
		return removeDuplicates(first.Next)
	}

	newHead := first

	second := first.Next
	for second != nil && second.Next != nil {

		if second.Val == second.Next.Val {
			for second.Next != nil && second.Val == second.Next.Val {
				second = second.Next
				continue
			}
			// Remove duplicates
			second = second.Next
		}

		first.Next = second
		first = second
		second = second.Next
	}

	return newHead
}

// Recursion
func recRemoveDuplicates(head *datastruct.ListNode) *datastruct.ListNode {
	// 长度 <=1 的 list ，可以直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 要么 head 重复了，那就删除 head
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return recRemoveDuplicates(head.Next)
	}

	// 要么 head 不重复，递归处理 head 后面的节点
	head.Next = recRemoveDuplicates(head.Next)
	return head
}
