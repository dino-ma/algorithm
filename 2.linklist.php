<?php
/**
 * 小灰漫画算法 链表实现
 * 一个基础不好在努力学习的PHPer
 * <dino_ma@163.com>
 */

//创建一个链表
class LinkList
{
    public $data;
    public $next;
}

$linkList = new LinkList();
$linkList->next = null;
$linkList->data = null;
for ($i = 0; $i < 4; $i++) {
    $tmpLinkList = new LinkList();
    $tmpLinkList->data = $i;
    $tmpLinkList->next = $linkList->next;
    $linkList->next = $tmpLinkList;
}
echo '创建一个链表'.PHP_EOL;
print_r($linkList);


function printNode($pHead)
{
    if (is_null($pHead)) {
        return null;
    }
    while (!is_null($pHead)) {
        echo $pHead->data.PHP_EOL;
        $pHead = $pHead->next;
    }
}

echo '打印链表'.PHP_EOL;
printNode($linkList);


//插入一个节点（头/中/尾）
function insertItemByHead($pHead, $item)
{
    if (is_null($pHead) || is_null($item)) {
        return null;
    }
    $item->next = $pHead->next;
    $new = new LinkList();
    $new->next = $item;

    return $new;
}
$item = new LinkList();
$item->data = 'tmpITEM';
$item->next = null;
echo '通过头节点插入一个元素';

$linkList = insertItemByHead($linkList, $item);
print_r($linkList);

//删除第N个节点
function delItemByHead($pHead, &$item)
{
    if (is_null($pHead)) {
        return null;
    }
    $head = $pHead->next;
    $flag = false;
    while (!is_null($head->next)) {
        if ($head->next->data == $item->data) {
            $flag = true;
            break;
        } else {
            $head = $head->next;
        }
    }
    if ($flag) {
        $head->next = $head->next->next;
    }
    return $pHead->next;
}

echo '删除某个节点'.PHP_EOL;
$item = new LinkList();
$item->data = 3;
$item->next = null;
print_r(delItemByHead($linkList, $item));

//查找一个节点
function findItemByLinkList($pHead, $item)
{
    if (is_null($pHead) || is_null($item)) {
        return null;
    }
    $head = $pHead->next;
    while (!is_null($head)) {
        if ($head->data == $item->data) {
            return $item->data;
        }
        $head = $head->next;
    }

     return false;
}
echo '查找某个节点并返回这个节点'.PHP_EOL;
$item = new LinkList();
$item->data = 'tmpITEM';
$item->next = null;
print_r(findItemByLinkList($linkList, $item));




//一个曾经面试遇到过的一个问题  查找链表中倒数第N个元素 最优解 链表无限大 N也可能无限大
function findNItemByLinkList($linkList, int $n)
{
    if (is_null($linkList) || $n <= 0) {
        return null;
    }
    $pre = $linkList;
    $cur = $linkList;
    $i = 1;
    while (!is_null($cur->next)) {
        $cur = $cur->next;
        if ($i>$n) {
            $pre = $pre->next;
        }
        ++$i;
    }

    $node = $pre->next;

    return $node->data;
}

$linkList = new LinkList();
$linkList->next = null;
$linkList->data = null;
for ($i = 0; $i < 4; $i++) {
    $tmpLinkList = new LinkList();
    $tmpLinkList->data = $i;
    $tmpLinkList->next = $linkList->next;
    $linkList->next = $tmpLinkList;
}
echo PHP_EOL.'查找倒数第N个元素'.PHP_EOL;

print_r($linkList);

var_dump(findNItemByLinkList($linkList, 1));


//链表反转
function reversList($pHead)
{
    if (is_null($pHead)) {
        return null;
    }
    $old = $pHead->next;
    $new = $tmp = null;
    while (!is_null($old)) {
        $tmp = $old->next;
        $old->next = $new;
        $new = $old;
        $old = $tmp;
    }
    $newHead = new LinkList();
    $newHead->next = $new;
    return $newHead;
}

echo '链表反转'.PHP_EOL;

print_r(reversList($linkList));
