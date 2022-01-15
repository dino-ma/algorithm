package main

import "fmt"

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2, 5, 5}))
}
