package main

type LRUCache struct {
	currentCap int
	cap        int
	hashmap    map[int]*LRUNode
	head       *LRUNode
	tail       *LRUNode
}

type LRUNode struct {
	key  int
	val  int
	next *LRUNode
	prev *LRUNode
}

func Constructor(capacity int) LRUCache {
	lruCache := new(LRUCache)
	lruCache.cap = capacity
	lruCache.hashmap = make(map[int]*LRUNode)
	lruCache.head = initNode(0, 0)
	lruCache.tail = initNode(0, 0)
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return *lruCache
}

func (this *LRUCache) Get(key int) int {
	var node *LRUNode
	var ok bool
	if node, ok = this.hashmap[key]; !ok {
		return -1
	}
	moveToHead(this.head, node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	var node *LRUNode
	var ok bool
	if node, ok = this.hashmap[key]; ok {
		moveToHead(this.head, node)
		node.val = value
		return
	}
	node = initNode(key, value)
	this.hashmap[key] = node
	addToHead(this.head, node)
	this.currentCap++
	if this.currentCap > this.cap {
		delete(this.hashmap, this.tail.prev.key)
		removeNode(this.tail.prev)
	}
}

func initNode(key int, val int) *LRUNode {
	lruNode := new(LRUNode)
	lruNode.key = key
	lruNode.val = val
	return lruNode
}

func moveToHead(head *LRUNode, node *LRUNode) {
	removeNode(node)
	addToHead(head, node)
}

func addToHead(head *LRUNode, node *LRUNode) {
	node.next = head.next
	head.next.prev = node
	head.next = node
	node.prev = head
}

func removeNode(currNode *LRUNode) {
	currNode.prev.next = currNode.next
	currNode.next.prev = currNode.prev
}

func main() {
	lruCache := Constructor(2)
	lruCache.Put(2, 1) // 缓存是 {2=1}
	lruCache.Put(1, 1) // 缓存是 {2=1, 1=1}
	lruCache.Put(2, 3) // 缓存是 {2=3, 1=1}
	lruCache.Put(4, 1) // 缓存是 {4=1, 2=3}
	lruCache.Get(1)    // 返回 -1 (未找到)
	lruCache.Get(2)    // 返回
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
