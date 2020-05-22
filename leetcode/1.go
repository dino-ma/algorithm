package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numsLength := len(nums);
	if numsLength < 1 {
		return nil;
	}
	hashMap := make(map[int]int);
	for numsKey, numsVal := range nums {

		hashMapVal, ok := hashMap[target - numsVal];
		if ok {
			return []int{numsKey, hashMapVal};

		}

		hashMap[numsVal] = numsKey;
	}

	return nil;
}

func main() {
	sum := 22
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, sum))
}
