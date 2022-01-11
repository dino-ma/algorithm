package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		pre = findCommonPrefix(pre, strs[i])
		if len(pre) == 0 {
			break
		}
	}
	return pre
}

// findCommonPrefix findCommonPrefix
func findCommonPrefix(str1, str2 string) string {
	strLen := findCommonPrefixMin(len(str1), len(str2))
	index := 0
	for index < strLen && str1[index] == str2[index] {
		index++
	}
	return str2[:index]
}

func findCommonPrefixMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
