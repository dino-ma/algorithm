package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root = new(TreeNode)
	root.Val = 2

	var childLeft = new(TreeNode)
	childLeft.Val = 1
	root.Left = childLeft

	var childRight = new(TreeNode)
	childRight.Val = 3
	root.Right = childRight
	fmt.Println(root)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)


	var otherRoot = new(TreeNode)
	otherRoot.Val = 5

	var otherChildLeft = new(TreeNode)
	otherChildLeft.Val = 1
	otherRoot.Left = otherChildLeft

	var otherChildRight = new(TreeNode)
	otherChildRight.Val = 4
	otherRoot.Right = otherChildRight


	var otherChildLeftChild = new (TreeNode)
	otherChildLeftChild.Val = 3
	otherChildRight.Left = otherChildLeftChild

	var otherChildRightChild = new (TreeNode)
	otherChildRightChild.Val = 6
	otherChildRight.Right = otherChildRightChild
	fmt.Println(otherRoot.Right.Right.Val)


}
