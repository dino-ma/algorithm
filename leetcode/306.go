package main

import (
	"fmt"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 || len(num) > 35 {
		return false
	}
	if num[0] == 0 {
		return false
	}
	data := strings.Split(num, "")
	var pre string
	pre = data[0]
	var prePre string
	prePre = data[1]
	for k, v := range data {
		if k == 0 || k == 1 {
			continue
		}
		result := pre + prePre
		pre = v
		fmt.Println(result)
	}
	return true
}

func main() {
	// fmt.Println("isAdditiveNumber:", isAdditiveNumber("1"))
	fmt.Println("isAdditiveNumber:", isAdditiveNumber("112"))
	// fmt.Println("isAdditiveNumber:", isAdditiveNumber("199100199"))
	// fmt.Println("isAdditiveNumber:", isAdditiveNumber("1235813"))

}
