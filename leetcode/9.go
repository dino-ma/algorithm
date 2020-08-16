package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	maxLength := len(str)
	if  maxLength == 0{
		return false
	}
	i := 0
	right := maxLength-1

	for i < right {
		if str[i] != str[right] {
			return false
		}
		i++
		right--
	}

	return true
}

func main()  {

	fmt.Println(isPalindrome(123456))
	fmt.Println(isPalindrome(12))
	fmt.Println(isPalindrome(123321))
}
