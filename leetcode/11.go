package main

import "fmt"

func Max(a,b int) int {
	if a>b {
		return a
	}

	return b
}

func Min(a,b int) int {
	if a< b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	length := len(height)
	maxArea := 0
	if length == 0{
		return maxArea
	}
	l := 0
	r := length - 1
	for l < r {
		left , right := height[l], height[r]
		area := (r-l) * Min(left, right)

		maxArea = Max(area, maxArea)
		if left > right {
			r--
		}else{
			l++
		}

	}

	return maxArea

}

func main()  {
	maxArea :=maxArea([]int{1,8,6,2,5,4,8,3,7})

	fmt.Println(maxArea)
}
