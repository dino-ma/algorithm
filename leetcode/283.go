package main

import "fmt"

func moveZeroes(nums []int) []int{
	length := len(nums)
	if length == 0{
		return nums
	}
	left := 0

	for i:=0; i<length; i++ {
		if nums[i] == 0 {
			continue
		}
		nums[left] = nums[i]
		if left != i {
			nums[i] = 0
		}
		left++

	}

	return nums
}

func main()  {
	nums := []int{0,1,0,3,12}
	//nums := []int{1}
	fmt.Println(moveZeroes(nums))
}
