package main

import "fmt"

func maxDepth(s string) int {
	var deep int
	if len(s) == 0 {
		return deep
	}
	var max int
	for _, v := range s {
		if v == '(' {
			deep++
			if deep >= max {
				max = deep
			}
		} else if v == ')' {
			deep--
		}
	}
	return max
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
	fmt.Println(maxDepth("1+(2*3)/(2-1)"))
}
