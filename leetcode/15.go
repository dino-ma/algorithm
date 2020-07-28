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
	sort.Ints(nums)//升序做排序 -4 -1 0 123
	fmt.Println(nums)
	result := make([][]int, 0, 0)//构造返回数据 默认 LEN CAP为0 依靠APPEND做扩容
	fmt.Println(length)
	//length - 1为了i和right不相交
	for i := 0; i < length-1; i++ {
		left, right := i+1, length-1
		//nums[i]> 0 代表 三数之和必大于0 ，或者 i 和 i+1大于0 则可以跳过本次循环
		if nums[i] > 0 || (nums[i]+nums[left]) > 0 {
			continue
		}
		//寻找三个指针 的相邻的值如果相等则跳过 避免重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//left right 不相交
		for left < right {
			//如果left指针 已跳过首次循环 则去除左侧相邻的相同值
			if left > i + 1 && nums[left] == nums[left-1] {
				left++
				continue

			}
			//如果right指针 已跳过首次循环 则去除左侧相邻的相同值
			if right < length - 2 && nums[right] == nums[right+1]{
				right--
				continue
			}

			sum := nums[i] + nums[left] + nums[right]
			//常规移动指针
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}else if sum > 0 {
				right--
			}else if sum < 0 {
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
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
