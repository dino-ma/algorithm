package main

import "fmt"

func isAnagram(s string, t string) bool {
	sL := len(s)
	tl := len(t)
	if sL != tl {
		return false
	}
	sArray := [26]int{}
	tArray := [26]int{}
	for _,v := range s  {
		sArray['v'-v]++
	}

	for _,childT := range t  {
		tArray['v'-childT]++
	}
	fmt.Println(sArray,tArray)

	return sArray == tArray
}


func main()  {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}