<?php

// +----------------------------------------------------------------------------
// | nice版权所属 
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 http://oneniceapp.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);

$string = ')))(((';
$string2 = '()][';
$string3 = '{()}';
$string4 = '[]()[](())[[]][[';
$string5 = '()';
function isTag(string $s) : bool
{
    $map = [
        ")" => "(",
        "}" => "{",
        "]" => "[",
    ];

    $len = strlen($s);
    if ( $len % 2 == 1) {
        return false;
    }
    $stack = [];

    for ($i = 0; $i < $len; $i++) {
        if (isset($map[$s[$i]])) {
            if (isset($stack) && $stack[0] == $map[$s[$i]]) {
                array_shift($stack);
            }else {
                return false;
            }
        } else{
            array_unshift($stack, $s[$i]);
        }
    }
    if (empty($stack)) {
        return true;
    }

    return false;
}


var_dump(isTag('()'));