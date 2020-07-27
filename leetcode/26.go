package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2{
		 return length
	}
	left , right := 0, 1
	for right < length {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
		right++
	}

	return left + 1
}

func main()  {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	len := removeDuplicates(nums)
	fmt.Println(len)
}
