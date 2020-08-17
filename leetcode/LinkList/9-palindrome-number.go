package main

import "strconv"

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	strLength := len(xStr)
	if strLength == 0 {
		return true
	}

	for i := 0; i< strLength ; i++  {
		if xStr[i] != xStr[strLength-1] {
			return false
		}
		strLength--
	}

	return true
}


