package main

import (
	"fmt"
	"math"
)

var maxNum = 3

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < maxNum {
		return false
	}
	first, second := nums[0], math.MaxInt64
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(increasingTriplet([]int{1, 3, 3, 4, 5, 6}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1, 0}))
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	fmt.Println(increasingTriplet([]int{6, 7, 1, 2}))
}
