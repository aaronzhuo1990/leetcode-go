package main

import (
	"fmt"
	"github.com/azhuox/leetcode-go/000datastruct"
)

func main() {
	testTree := datastruct.PreIn2Tree([]int{1, 2, 3}, []int{1, 3, 2})
	fmt.Printf("result %v\n", inorderTraversal(testTree)) // Should be 1, 3, 2
}

func inorderTraversal(root *datastruct.BinaryTreeNode) []int {
	// Basic case
	if root == nil {
		return nil
	}

	// Recursive case
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}
