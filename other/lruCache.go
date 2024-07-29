package main

type LRUNode struct {
	key  int
	val  int
	prev *LRUNode
	next *LRUNode
}

type LRUCache struct {
	nowCap int
	cap    int
	hash   map[int]*LRUNode
	head   *LRUNode
	tail   *LRUNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &LRUNode{}, &LRUNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{nowCap: 0, cap: capacity, hash: map[int]*LRUNode{}, head: head, tail: tail}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hash[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.hash[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}
	newNode := &LRUNode{key: key, val: value}
	this.addToHead(newNode)
	this.hash[key] = newNode
	this.nowCap++
	if this.nowCap > this.cap {
		this.removeTail()
		delete(this.hash, this.tail.prev.key)
		this.nowCap--
	}
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
	node.next.prev = node
}

func (this *LRUCache) removeNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) removeTail() {
	this.removeNode(this.tail.prev)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

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
