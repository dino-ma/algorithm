<?php
/**
 * 小灰漫画算法 数组实现
 * 一个基础不好在努力学习的PHPer
 * <dino_ma@163.com>
 */

$array = [1,2,3,4,5];

//中间插法
function insertMiddleArray(int $cap, string $val, array $array) : array
{
    if (empty($array)) {
        return [];
    }

    $size = count($array);

    if ($cap < 0 || $cap > $size) {
        return $array;
    }
    //中间插讲究的是倒叙开始，并当目标index为移动的最后一个index后则循环结束
    for ($i = $size - 1; $i >= $cap; $i-- )
    {
        $array[$i+1] = $array[$i];
    }
    $array[$cap] = $val;

    return $array;
}

echo '中间插法'.PHP_EOL;
print_r(insertMiddleArray(2, 'a', $array));

//超范围插入数组
function overRangeInsertArray(int $index, string $val, array $array) : array
{
    if (empty($array)) {
        return [];
    }

    if ($index < 0) {
        return [];
    }
    $size = count($array);
    if ($index > $size) {
        $newSize = $index-$size;
        //数组扩容仅扩容需要的空间
        $array = array_merge($array, array_fill($size, $newSize, ''));
        $size = $newSize;
    }

    for ($i = $size - 1; $i >= $index; $i--) {
        $array[$i+1] = $array[$i];
    }
    $array[$index] = $val;

    return $array;
}

echo '超范围插入法'.PHP_EOL;
print_r(overRangeInsertArray(9, 'a', $array));

function deleteItem(int $index, array $array) : array
{
    if (empty($array)) {
        return [];
    }
    $size = count($array);
    if ($index > $size) {
        //溢出不处理
        return $array;
    }

    for ($i = $index ; $i < $size; $i++) {
        $array[$i] = $array[$i+1];
    }

    return $array;
}
echo '删除指定元素'.PHP_EOL;
print_r(deleteItem(1, $array));

//数组的优势和劣势
//数组拥有非常高效的随机访问能力 例如二分查找就利用了数组的这种优势
//数组的劣势
//数组的元素连续紧密存储在内存中，插入删除会导致大量的元素被迫移动影响效率

//适合读操作多，写操作少的场景。