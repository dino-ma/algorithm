package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root = new(TreeNode)
	root.Val = 0

	var childLeft = new(TreeNode)
	childLeft.Val = 1

	var childRight = new(TreeNode)
	childRight.Val = 2

	root.Left = childLeft
	root.Right = childRight

	fmt.Println(root)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)

}
