package main

import "math"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	return dfs(root,math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, min,max int) bool {

	if root == nil {
		return true
	}

	if min >= root.Val || max <= root.Val {
		return false
	}

	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}



func main()  {
	
}
