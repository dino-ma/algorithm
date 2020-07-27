package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0, 0)
	for i := 0; i < length-1; i++ {
		l := i + 1
		r := length - 1
		//去掉不可能的数
		if nums[i] > 0 || nums[i]+nums[l] > 0 {
			break
		}
		//跳过左右边界的重复项
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for l < r {
			if l > i + 1 && nums[l] == nums[l-1] {
				l++
				continue
			}

			if r < length - 2  && nums[r] == nums[r+1]{
				r--
				continue
			}
			x:= nums[i] + nums[l] + nums[r]
			if x == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if x < 0 {
				l++
			}else if x > 0 {
				r--
			}
		}
	}

	return result
}

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}x

	fmt.Println(threeSum(nums))
}
