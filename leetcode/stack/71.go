package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	if len(path) < 1 || len(path) > 3000 {
		return path
	}
	// stack 初始化栈
	stack := make([]string, 0, len(path)%3)
	// 将path用/分割为富足
	for _, v := range strings.Split(path, "/") {
		// 如果是 .. 则出栈，并且此处需要优先，否则会先入栈
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			// 空和单独的.可以过滤，又不是.. 又不是空 还不是.的则可以入栈
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("./././././././././"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
}
