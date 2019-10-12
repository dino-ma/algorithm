<?php
// +----------------------------------------------------------------------------
// | nice版权所属
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 http://oneniceapp.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);

class Queue
{

    protected $length;

    protected $maxLength;

    protected $popData;

    protected $error;

    protected $data;

    protected $outItem;

    public function getError() : string
    {
        return $this->error;
    }

    public function setMaxLength(int $length) : bool
    {
        if (!is_null($this->maxLength)) {
            $this->error = 'is seted';
            return false;
        }
        $this->maxLength = $length;
        return true;
    }

    public function getMaxLength() : int
    {
        return (int)$this->maxLength;
    }


    public function getLength() : int
    {
        return (int)$this->length;
    }

    public function push(string $item) : bool
    {
        if ($this->length + 1 > $this->maxLength) {
            $this->length = 'this queue is more than maxLength';
            return false;
        }
        ++$this->length;
        $this->data[$this->length] = $item;

        return true;
    }

    public function pop() : bool
    {
        if ($this->getLength() <= 0) {
            $this->error = 'this queue is empty';
            return false;
        }
        --$this->length;
        $this->outItem =  array_shift($this->data);

        return true;
    }

    public function getOutItem()
    {
        return $this->outItem;
    }

    public function getQueue() : array
    {
        return $this->data;
    }
}

$queueObj = new Queue();
$queueObj->setMaxLength(10);
$queueObj->push('2');
$queueObj->push('1');

print_r($queueObj->getQueue());
$queueObj->pop();
print_r($queueObj->getOutItem());