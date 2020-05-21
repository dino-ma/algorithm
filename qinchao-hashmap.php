<?php

// +----------------------------------------------------------------------------
// | nice版权所属 
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 http://oneniceapp.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);


function isSame(string $s, string $t) : bool
{
    $stringANum = strlen($s);
    $stringBNum = strlen($t);
    if ($stringANum != $stringBNum) {
        return false;
    }

    $stringMapA = [];
    $stringMapB = [];
    for($i = 0; $i < $stringANum; $i++) {
        if (isset($stringMapA[$s[$i].'-a'])) {
            $stringMapA[$s[$i].'-a']++;
        } else {
            $stringMapA[$s[$i].'-a'] = 1;
        }
    }

    for($i = 0; $i < $stringBNum; $i++) {
        if (isset($stringMapB[$t[$i].'-a'])) {
            $stringMapB[$t[$i].'-a']++;
        } else {
            $stringMapB[$t[$i].'-a'] = 1;
        }
    }
    return $stringMapB == $stringMapA;
}



$stringA = 'anagram';
$stringB = 'nagaram';

var_dump(isSame($stringA, $stringB));







