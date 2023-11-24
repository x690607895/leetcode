package main

// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

// 核心思想
// cap表示容量，len表示
// keys存储缓存的数据，使用map，get和set速度快
// 双向链表-》队列，每次删除都删除头部指针，每次插入在末尾

type LRUCache struct {
	capacity   int
	keys       map[int]*NodeData
	head, tail *NodeData
}

type NodeData struct {
	next, prev *NodeData
	K, V       int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keys:     make(map[int]*NodeData),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.keys[key]; ok {
		this.Remove(value)
		this.Add(value)
		return value.V
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if oldValue, ok := this.keys[key]; ok {
		oldValue.V = value
		this.Remove(oldValue)
		this.Add(oldValue)
	} else {
		nodeData := &NodeData{K: key, V: value}
		this.Add(nodeData)
		this.keys[key] = nodeData
		if len(this.keys) > this.capacity {
			delete(this.keys, this.head.K)
			this.Remove(this.head)
		}
	}
}

func (this *LRUCache) Remove(nd *NodeData) {
	if this.head == nd {
		this.head = nd.next
		if nd.next != nil {
			nd.next.prev = nil
		}
	} else if this.tail == nd {
		this.tail = nd.prev
		if nd.prev != nil {
			nd.prev.next = nil
		}
	} else {
		nd.prev.next = nd.next
		nd.next.prev = nd.prev
	}
}

func (this *LRUCache) Add(nd *NodeData) {
	if this.head == nil {
		this.head = nd
	} else {
		nd.prev = this.tail
		this.tail.next = nd
	}
	this.tail = nd
}
