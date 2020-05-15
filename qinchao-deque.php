<?php

// +----------------------------------------------------------------------------
// | nice版权所属 
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 http://oneniceapp.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);


class Solution{

    /**
     * @param array $nums
     * @param int $k    窗口长度
     * @return intger []
     * @throws Exception
     */
    function maxSlidingWindows(array $nums, int $k)
    {
        $count = count($nums);
        if ($count < 2 || $k > $count || $k < 1) {
            return $nums;
        }
        $window = [];
        $res = [];

        for ($i = 0; $i < $count; $i++) {
            //构建窗口后 移动窗口时 弹出队列左侧元素写入新的队列元素
            if ($i >= $k + $window[0]) {
                array_shift($window);
            }
            //循环比较 保证最右的边界是最小的
            while ($window && $nums[end($window)] <= $nums[$i]) {
                array_pop($window);
            }
            //构造窗口
            $window[] = $i;
            //移动过了一个窗口则需要取窗口中目前的最左侧的值作为结果
            if ($i >= $k - 1) {
                $res[] = $nums[$window[0]];
            }
        }

        return $res;
    }
}

$obj = new Solution();
$a = $obj->maxSlidingWindows([1,3,-1,-3,5,3,6,7], 3);
print_r($a);