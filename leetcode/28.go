package main

import "fmt"

func strStr(haystack string, needle string) int {
	hayLen := len(haystack)
	needLen := len(needle)
	if hayLen < needLen {
		return -1
	}
	// 计算窗口即可
	// 细节点起始可以为needLen而不用一定从0开始
	for i := needLen; i <= hayLen; i++ {
		if haystack[i-needLen:i] == needle {
			return i - needLen
		}
	}
	return -1
}
func main() {
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("hello", "e"))
	fmt.Println(strStr("hello", "h"))
	fmt.Println(strStr("aaaaaa", "bba"))
}
