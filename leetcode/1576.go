package main

import "fmt"

func modifyString(s string) string {
	var l int
	l = len(s)
	if l == 0 {
		return ""
	}
	strByte := []byte(s)
	for i, v := range strByte {
		if v != '?' {
			continue
		}
		// 此处循环从a开始到C结束，因为只需要不连续重复即可，那么则为 2+1
		for word := byte('a'); word <= byte('c'); word++ {
			// 控制边界，问号可能出现在第一个和最后一个
			if !(i > 0 && strByte[i-1] == word || i < l-1 && strByte[i+1] == word) {
				strByte[i] = word
				break
			}
		}

	}
	return string(strByte)
}

func main() {
	fmt.Println(modifyString("j?qg??b"))
}
