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
//插入一个节点（头/中/尾）

//删除一个节点

//查找一个节点

//输出链表
