package main

import "fmt"

func moveZeroes(nums []int)  []int{
	length := len(nums)
	if length == 0 {
		return nums
	}
	left := 0

	for i:=0; i<length; i++ {
		if nums[i] != 0 {
			nums[left] = nums[i]
			if i != left {
				nums[left] = 0
			}
			left++
		}
	}

	return nums

}

func main()  {

	result := moveZeroes([]int{0,1,0,3,12})
	fmt.Println(result)
}
