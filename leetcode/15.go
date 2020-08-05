package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	maxLength := len(nums)
	if maxLength < 3 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0, 0)
	for i := 0; i< maxLength-1 ; i++ {
		left ,right := i+1, maxLength-1
		if nums[i]>0 || nums[i] + nums[left] > 0  {
			continue
		}

		if i>0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			if left > i+1  && nums[left] == nums[left-1]{
				left++
				continue
			}

			if right < maxLength-2 && nums[right] == nums[right+1]{
				right--
				continue
			}


			x := nums[i] + nums[left] +nums[right]
			if x == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if x >0 {
				right--
			} else if x < 0 {
				left++
			}
		}
	}


	return result
}

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//2分sort O(Nlogn) find
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
}
