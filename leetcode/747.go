package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	m1, m2, idx := -1, -1, -1
	for k, v := range nums {
		if v > m1 {
			m1, m2, idx = v, m1, k
		} else if v > m2 {
			m2 = v
		}
	}
	if m1 >= m2*2 {
		return idx
	}
	return -1
}
func main() {
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}
