package main

import "fmt"

func twoSum(nums []int, target int) []int {
	maxLength := len(nums)
	if maxLength < 2 {
		return nil
	}
	hashMap := make(map[int]int)

	for i := 0; i < maxLength; i++ {
		hashVal , ok := hashMap[target-nums[i]]
		if ok {
			return []int{hashVal, i}
		}
		hashMap[nums[i]] = i
	}

	return nil
}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

}
