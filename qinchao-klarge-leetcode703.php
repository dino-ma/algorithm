<?php

// +----------------------------------------------------------------------------
// | nice版权所属 
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 http://oneniceapp.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);
//数据流中第K大的元素 leetcode 703

class KthLargest {

    private $objHeap; //小顶堆对象

    private $maxNum;//小顶堆最多成员

    /**
     * @param Integer $k
     * @param Integer[] $nums
     */
    function __construct($k, $nums) {
        if ($k <= 0) {
            return null;
        }
        if (!is_array($nums)) {
            return null;
        }
        $this->objHeap = new SplMinHeap();
        $this->maxNum = $k;
        $count = count($nums);
        for ($i = 0; $i < $count; $i++) {
            $this->add($nums[$i]);
        }

    }

    /**
     * @param Integer $val
     * @return Integer
     */
    function add($val) {
        if ($this->objHeap->count() < $this->maxNum) {
            $this->objHeap->insert($val);
        } else if ($val > $this->objHeap->top()) {
            $this->objHeap->extract();
            $this->objHeap->insert($val);
        }

        return $this->objHeap->top();
    }

}

/**
 * Your KthLargest object will be instantiated and called as such:
 * $obj = KthLargest($k, $nums);
 * $ret_1 = $obj->add($val);
 */

$obj = new KthLargest(3, [6,7]);
var_dump($obj->add(9));
var_dump($obj->add(2));
var_dump($obj->add(10));
var_dump($obj->add(18));