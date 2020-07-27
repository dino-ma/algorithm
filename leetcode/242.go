package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool  {
	sLen := len(s);
	tLen := len(t);
	if sLen != tLen{
		 return false;
	}
	sArray := [26]int{};
	tArray := [26]int{};
	for _, v := range s{
		sArray[v-'a']++;
	}
	for _, v := range t {
		tArray[v-'a']++;
	}

	return sArray == tArray;
}

//valid-anagram

func main()  {
	var s = "anagram";
	var t ="nagaram";
	fmt.Print(isAnagram(s, t));
}