package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	numsLength := len(nums)
	if numsLength == 0 || k <= 0 {
		return nil
	}
	queue := make([]int, 0, numsLength)
	result := make([]int, numsLength-k+1)
	for i := 0; i < numsLength; i++ {

		for len(queue) != 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		if queue[0] == i-k {
			queue = queue[1:]
		}

		resultKey := i + 1 - k
		if resultKey >= 0 {
			result[resultKey] = nums[queue[0]]
		}

	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}
