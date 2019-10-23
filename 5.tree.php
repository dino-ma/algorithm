<?php
// +----------------------------------------------------------------------------
// | nice版权所属
// +----------------------------------------------------------------------------
// | Copyright (c) 2019 http://oneniceapp.com All rights reserved.
// +----------------------------------------------------------------------------
// | Author: dino.ma <dino_ma@163.com>
// +----------------------------------------------------------------------------
declare(strict_types=1);

class BinaryNode
{
    private $left;
    private $right;
    private $data;

    public function __construct(string $data = '')
    {
        $this->data = $data;
    }

    public function addChildRen(BinaryNode $left, BinaryNode $right)
    {
        $this->left = $left;
        $this->right = $right;
    }

    public function getData() : string
    {
        return (string)$this->data;
    }

    public function getLeft()
    {
        return $this->left;
    }

    public function getRight()
    {
        return $this->right;
    }
}


class BinaryTree
{
    public $root = null;

    public function __construct(BinaryNode $node)
    {
        $this->root = $node;
    }

    public function isEmpty()
    {
        return empty($this->root);
    }

    public function printTree(BinaryNode $node, int $level = 0)
    {
        if ($node) {
            echo str_repeat('-', $level).$node->getData().PHP_EOL;

            if ($node->getLeft()) {
                $this->printTree($node->getLeft(), $level+1);
            }

            if ($node->getRight()) {
                $this->printTree($node->getRight(), $level+1);
            }
        }
    }
}


//构造一个二叉树
$root = new BinaryNode('root');
$tree = new BinaryTree($root);

$father1 = new BinaryNode('father1');
$father2 = new BinaryNode('father2');

$childRen1 = new BinaryNode('child1-2');
$childRen2 = new BinaryNode('child1-3');
$childRen3 = new BinaryNode('child1-2-1');
$childRen4 = new BinaryNode('child1-3-1');

$childRen2->addChildRen($childRen4, new BinaryNode());
$childRen1->addChildRen($childRen3, new BinaryNode());
$father1->addChildRen($childRen1, $childRen2);

$childRen21 = new BinaryNode('child2-2');
$childRen22 = new BinaryNode('child2-3');
$childRen23 = new BinaryNode('child2-2-1');
$childRen24 = new BinaryNode('child2-3-1');

$childRen21->addChildRen($childRen23, new BinaryNode());
$childRen22->addChildRen($childRen24, new BinaryNode());
$father2->addChildRen($childRen21, $childRen22);

$root->addChildRen($father1, $father2);
$tree->printTree($tree->root);

