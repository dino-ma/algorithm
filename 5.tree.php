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


//二叉搜索树 Binary Sort Tree
class BinarySortTree
{

    public $data;

    public $left;

    public $right;


    public function __construct()
    {
    }


    public function setData(int $data) : BinarySortTree
    {
        $this->data = $data;
        return $this;
    }


    /**
     * 构造一个二叉树
     * @param int $data
     * @param BinarySortTree $tree
     * @return BinarySortTree
     */
    public function insert(int $data, BinarySortTree $tree)
    {

        if (is_null($tree->data)) {
            $tree->data = $data;
            return $tree;
        }

        $node = $tree;

        while ($node) {
            if ($node->data > $data) {
                if ($node->left) {
                    $node = $node->left;
                } else {
                    $nodeObj = new BinarySortTree();
                    $nodeObj->data = $data;

                    $node->left = $nodeObj;
                    $node = $node->left;

                    break;
                }
            } elseif ($node->data < $data) {
                if ($node->right) {
                    $node = $node->right;
                } else {
                    $nodeObj = new BinarySortTree();
                    $nodeObj->data = $data;

                    $node->right = $nodeObj;
                    $node = $node->right;

                    break;
                }
            } else {
                break;
            }
        }

        return $node;
    }



    public function min(BinarySortTree $tree) : BinarySortTree
    {
        while ($tree->left) {
            $tree = $tree->left;
        }

        return $tree;
    }

    public function max(BinarySortTree $tree) : BinarySortTree
    {
        while ($tree->right) {
            $tree = $tree->right;
        }

        return $tree;
    }


    /**
     * 二叉树前序遍历
     * 根节点->左子树->右子树
     * @param BinarySortTree $tree
     * @return array
     */
    public function preOrderTraversal(BinarySortTree $tree) : array
    {
        if (is_null($tree)) {
            return [];
        }
        $traversal = $data = [];
        array_push($traversal, $tree);

        while (!empty($traversal)) {
            $item = array_pop($traversal);
            $data[] = $item->data;

            if (!is_null($item->right)) {
                array_push($traversal, $item->right);
            }
            if (!is_null($item->left)) {
                array_push($traversal, $item->left);
            }
        }

        return $data;
    }


    /**
     * 中序遍历
     * 左子树->根节点->右子树
     * @param BinarySortTree $tree
     * @return array
     */
    public function sequentialTraversal(BinarySortTree $tree) : array
    {
        $stack = array();
        $current_node = $tree;
        while (!empty($stack) || $current_node != null) {
            while ($current_node != null) {
                array_push($stack, $current_node);
                $current_node = $current_node->left;
            }
            $current_node = array_pop($stack);
            echo $current_node->data . " ";
            $current_node = $current_node->right;
        }

        return [];
    }


    /**
     * 后序遍历
     * 左子树->右子树->跟节点
     * @param BinarySortTree $tree
     * @return array
     */
    public function postOrderTraversal(BinarySortTree $tree) : array
    {

    }




}

//构造一个二叉搜索树
$node = new BinarySortTree();
$root = $node->insert(500, $node);
$node->insert(499, $root);
$node->insert(496, $root);
$node->insert(495, $root);

$node->insert(501, $root);
$node->insert(498, $root);
$node->insert(497, $root);
$node->insert(503, $root);
$node->insert(502, $root);
$node->insert(505, $root);
print_r($root);

$min = $node->min($root);
print_r($min);
$max = $node->max($root);
print_r($max);
//AVL平衡树、红黑树、树堆等

//遍历
//深度优先： 前序、中序、后序
//广度优先：层序


//前序遍历
$traversal = $node->preOrderTraversal($root);
print_r($traversal);//500 499 496 495 498 497 501 503 502 505

//中序遍历
$traversal = $node->sequentialTraversal($root);
print_r($traversal);//497 495 496 498 499 500 501 503 502 505
