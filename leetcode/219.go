package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := map[int]bool{}
	for i, v := range nums {
		if i > k {
			delete(hash, nums[i-k-1])
		}
		if _, ok := hash[v]; ok {
			return true
		}
		hash[v] = true
	}
	return false
}
func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

}
