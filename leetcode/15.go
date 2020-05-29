package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	numsLength := len(nums)
	if numsLength < 3 {
		return [][]int{}
	}
	//排序 为了后面的去重和双指针做准备
	sort.Ints(nums)
	//这个时候数据已经是 -4 -1 -1 0 1 2
	fmt.Println(numsLength)
	resultArray := make([][]int, 0, 0) //声明结果数组

	for i := 0; i < numsLength-1; i++ {
		l, r := i+1, numsLength-1
		if nums[i] > 0 || nums[i]+nums[l] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} // i 去重

		fmt.Println(l, r)

		//l和r永远不相交
		for l < r {
			//去掉左侧有重复的值
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			//去掉右侧重复的值

			if r < numsLength-2 && nums[r] == nums[r+1] {
				r--
				continue
			}

			x := nums[i] + nums[l] + nums[r]

			if x == 0 {
				resultArray = append(resultArray, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if x > 0 {
				r--
			} else if x < 0 {
				l++
			}
		}
	}

	return resultArray
}

func main() {
	//2分sort O(Nlogn) find
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
}
