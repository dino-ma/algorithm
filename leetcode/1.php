<?php

// +----------------------------------------------------------------------------
// | nice版权所属 
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 http://oneniceapp.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);



class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {

        $length = count($nums);
        $map = [];
        for ($i = 0; $i < $length; $i++) {
            if (isset($map[$target - $nums[$i]])) {
                return [$i, $map[$target - $nums[$i]]];
            }
            $map[$nums[$i]] = $i;
        }


        return null;
    }
}



var_dump((new Solution())->twoSum([2, 7, 11, 15], 9));