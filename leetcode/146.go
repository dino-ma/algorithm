package main


type LRUCache struct {
	cap int
	cache map[int]*DListLinkNode
	head , tail *DListLinkNode
}

//构造双向链表
type DListLinkNode struct {
	key , value int
	pre , next *DListLinkNode
}

//添加节点
func (this *LRUCache)AddNode(node *DListLinkNode)  {
	head := this.head
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node

}

//删除节点
func (this *LRUCache)RemoveNode(node *DListLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//移动某个节点到头部
func (this *LRUCache)MoveNodeToHead(node *DListLinkNode)  {
	this.RemoveNode(node)
	this.AddNode(node)
}

func Constructor(capacity int) LRUCache {
	head :=&DListLinkNode{0,0,nil,nil}
	tail :=&DListLinkNode{0,0,nil,nil}
	head.next = tail
	tail.pre = head
	return LRUCache{
		cap:capacity,
		cache: make(map[int]*DListLinkNode),
		head:head,
		tail:tail,
	}
}


func (this *LRUCache) Get(key int) int {
	node , exists := this.cache[key]
	if !exists {
		return -1
	}
	this.MoveNodeToHead(node)

	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	node , exists := this.cache[key]
	tail := this.tail
	if exists {
		//赋值
		node.value = value
		//移动链表
		this.MoveNodeToHead(node)
	} else {
		node := &DListLinkNode{key:key,value:value, pre:nil, next:nil}
		//看容量是否足够
		if len(this.cache) >= this.cap {
			delete(this.cache, tail.pre.key)
			this.RemoveNode(tail.pre)
		}
		this.AddNode(node)
		this.cache[key] = node
	}
}



//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得关键字 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得关键字 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
func main()  {
	//LRU缓存最少使用原则 我需要设置一个缓存的容量

	//构造一个双向链表

	//插入一个节点 从头部插入

	//删除某个节点 改变前驱和后继

	//移动节点到头节点

	//判断在CACHE中则更新并且移动到头部

	//如果不在CACHE中则判断容量，删除末尾节点中的上一个 tail->pre

}