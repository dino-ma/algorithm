package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	maxLength := len(height)
	maxArea := 0 //最大的面积
	if maxLength < 2 {
		return maxArea
	}
	left, right := 0, maxLength-1
	for left < right {
		length := right - left
		maxArea = max(maxArea, min(height[left], height[right])*length)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器，且 n 的值至少为 2。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	maxArea := maxArea([]int{2, 3, 4, 5, 18, 17, 6})
	fmt.Println(maxArea)
}
