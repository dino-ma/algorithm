<?php
// +----------------------------------------------------------------------------
// | dino.ma版权所属
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 https://mashengjie.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);

class Stack
{

    protected $top;

    protected $size;

    protected $maxSize;

    protected $error;

    protected $data;

    protected $popItem;


    public function setMaxSize(int $maxSize) : bool
    {
        if (!is_null($this->maxSize)) {
            $this->error = 'not repeat set maxSize';
            return false;
        }

        $this->maxSize = $maxSize;
        return true;
    }

    public function push(string $item) : bool
    {
        if ($this->top >= $this->maxSize) {
            $this->error = 'top is gt maxsize';
            return false;
        }
        ++$this->top;
        $this->data[$this->top] = $item;
        return true;
    }

    public function pop() : bool
    {
        //空栈返回false
        if ($this->top <= 0) {
            $this->error = 'top is elt 0';
            return false;
        }
        $this->popItem = $this->data[$this->top];
        unset($this->data[$this->top]);
        $this->top--;

        return true;
    }

    public function getPopItem() : string
    {
        return $this->popItem;
    }

    public function getError() : string
    {
        return $this->error;
    }
}

$stockObj = new Stack();
$stockObj->setMaxSize(10);
for ($i = 0; $i <= 5; $i++) {
    $stockObj->push((string)$i);
}
$stockObj->pop();

print_r($stockObj);
echo $stockObj->getPopItem();
