package datastruct

import "fmt"

// BinaryTreeNode Definition for a binary tree node.
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	panic(fmt.Errorf("the value %d nont fount in %v", val, nums))
}

// Convert the given preoder & inorder to a binary tree
func PreIn2Tree(pre, in []int) *BinaryTreeNode {
	if len(pre) != len(in) {
		panic(fmt.Errorf("the length of pre order %v does not equal to the length of %v", pre, in))
	}

	if len(in) == 0 {
		return nil
	}

	res := &BinaryTreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}
