package problem0146

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Key   int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	capcity int
	cache   map[int]*Node
	tail    *Node
	head    *Node
}

func (this *Node) ToString() string {
	var s strings.Builder
	tail := this
	for tail != nil {
		s.WriteString(fmt.Sprintf("(%d, %d) -> ", tail.Key, tail.Value))
		tail = tail.Next
	}

	return s.String()
}

func (this *Node) ToStringReverse() string {
	var s strings.Builder
	head := this
	for head != nil {
		s.WriteString(fmt.Sprintf("(%d, %d) <- ", head.Key, head.Value))
		head = head.Prev
	}

	return s.String()
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capcity: capacity, cache: make(map[int]*Node), head: nil, tail: nil}
}

func (this *LRUCache) PrintCacheKey() {
	for k, _ := range this.cache {
		fmt.Printf("Cache key: %d, ", k)
	}
	fmt.Println()
}

func (this *LRUCache) PrintList() {
	fmt.Printf("Head key %d, value: %d\n", this.head.Key, this.head.Value)
	fmt.Printf("Tail key %d, value: %d\n", this.tail.Key, this.tail.Value)
	fmt.Printf("List: %v\n", this.tail.ToString())
	fmt.Printf("List Reverse: %v\n", this.head.ToStringReverse())
}

func (this *LRUCache) addToNode(key int, value int) {
	// Handle if the current key is already set
	if this.cache[key] != nil {
		this.cache[key].Value = value	
		this.moveToHead(key)
		return
	}

	this.cache[key] = &Node{Value: value, Key: key, Prev: this.head, Next: nil}
	if len(this.cache) == 1 {
		this.tail = this.cache[key]
	}
	if this.head != nil {
		this.head.Next = this.cache[key]
	}
	this.head = this.cache[key]
}

func (this *LRUCache) moveToHead(key int) {
	if this.head == this.cache[key] {
		return
	}

	if this.tail == this.cache[key] {
		// Update tail to next to tail
		this.tail = this.cache[key].Next
		this.tail.Prev = nil
		// Update head with current cache[key]
		this.head.Next = this.cache[key]
		this.cache[key].Prev = this.head
		this.cache[key].Next = nil
		this.head = this.cache[key]
		return
	}

	this.cache[key].Prev.Next = this.cache[key].Next
	this.cache[key].Next.Prev = this.cache[key].Prev

	this.head.Next = this.cache[key]
	this.cache[key].Prev = this.head
	this.cache[key].Next = nil
	this.head = this.cache[key]
}

func (this *LRUCache) removeTailNode() {
	delete(this.cache, this.tail.Key)
	nextFromTail := this.tail.Next
	nextFromTail.Prev = nil
	this.tail = nextFromTail
}

func (this *LRUCache) Get(key int) int {
	if this.cache[key] == nil {
		return -1
	}

	this.moveToHead(key)

	return this.cache[key].Value
}

func (this *LRUCache) Put(key int, value int) {
	this.addToNode(key, value)
	if len(this.cache) > this.capcity {
		this.removeTailNode()
	}
}
