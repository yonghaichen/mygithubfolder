package main

type LRUCache struct {
	capacity int
	len      int
	hashMap  map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);

 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Constructor(capacity int) LRUCache {
	m := make(map[int]*Node)
	lru := LRUCache{
		capacity: capacity,
		hashMap:  m,
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

func (cache *LRUCache) Get(key int) int {
	if v, ok := cache.hashMap[key]; ok {
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		n := cache.head.Next
		cache.head.Next = v
		v.Prev = cache.head
		n.Prev = v
		v.Next = n
		return v.Value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {

	if v, ok := cache.hashMap[key]; ok {
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		n := cache.head.Next
		cache.head.Next = v
		v.Prev = cache.head
		n.Prev = v
		v.Next = n
		v.Value = value
		return
	}

	if cache.len < cache.capacity {
		cache.len++
		node := &Node{
			Key:   key,
			Value: value,
		}
		cache.hashMap[key] = node
		n := cache.head.Next
		cache.head.Next = node
		node.Prev = cache.head
		node.Next = n
		n.Prev = node

	} else {
		t := cache.tail.Prev
		cache.tail.Prev.Prev.Next = cache.tail
		cache.tail.Prev = cache.tail.Prev.Prev
		t.Value = value
		delete(cache.hashMap, t.Key)
		t.Key = key
		cache.hashMap[key] = t
		hn := cache.head.Next
		cache.head.Next = t
		t.Prev = cache.head
		t.Next = hn
		hn.Prev = t
	}

	return
}
