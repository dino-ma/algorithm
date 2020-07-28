package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	length := len(nums)
	if length < 2{
		return nil
	}

	hashMap := make(map[int]int)
	for key,val := range nums {

		hashMapVal , ok := hashMap[target - val]
		if ok {
			return []int{hashMapVal, key}
		}
		hashMap[val] = key
	}

	return nil
}

func main() {
	sum := 22
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, sum))
}
