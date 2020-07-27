package main

import "fmt"

func getSliceNum( m int64,  n uint32,  num uint32) uint32 {

	if m < n {
		return num
	}

	for i:=1; i < m; i++ {

		//1 2 3 4 5
		//m5-1 = 4 n=2
		//m4-1 = 3 n=1
		//m3-3 = 0 n=0

		getSliceNum(m, n, num)
	}
	



	return num
}

func main() {
	//M长的线段 切成N段 有多少种切法

	m := 5
	n := 3

	fmt.Println(getSliceNum(m,n,0))




	//目标明确，价值 解决什么问题
	//表述重点
	//
}