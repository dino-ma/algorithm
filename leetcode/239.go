package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length < k || length == 0 {
		return nil
	}
	result := make([]int, length - k + 1 );
	queue := make([]int, 0, length);
	for i :=0; i<length ;i++  {
		//队列不为空并且当前的值大于队列最左侧的值则只保留队列最末尾的值
		if len(queue) !=0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		//队列从头入
		queue = append(queue, i);
		fmt.Println(queue)

		//窗口滑动
		if queue[0] == i - k {
			queue = queue[1:]
		}
		//记录每个窗口的结果 i+1 - k 从第一个形成为窗口开始记录
		if i+1-k >= 0 {
			result[i+1 - k ] = nums[queue[0]];
		}
	}

	return result;
}

func main() {
	nums := []int{1,3,-1,-3,5,3,6,7}
	fmt.Println(maxSlidingWindow(nums, 3))
}
