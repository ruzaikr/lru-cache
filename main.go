package main

import "log"

// Do not edit the class below except for the insertKeyValuePair,
// getValueFromKey, and getMostRecentKey methods. Feel free
// to add new properties and methods to the class.
type LRUCache struct {
	maxSize       int
	head          *DoublyLinkedListNode
	tail          *DoublyLinkedListNode
	keyToListNode map[string]*DoublyLinkedListNode
}

type DoublyLinkedListNode struct {
	key  string
	val  int
	prev *DoublyLinkedListNode
	next *DoublyLinkedListNode
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		maxSize:       size,
		head:          nil,
		tail:          nil,
		keyToListNode: map[string]*DoublyLinkedListNode{},
	}
}

func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	if cache.head == nil {
		var newNode = &DoublyLinkedListNode{
			key:  key,
			val:  value,
			prev: nil,
			next: nil,
		}
		cache.head = newNode
		cache.tail = newNode
		cache.keyToListNode[key] = newNode
		return
	}

	if duplicateNode, keyExists := cache.keyToListNode[key]; keyExists {
		duplicateNode.val = value
		return
	}

	var currentSize = len(cache.keyToListNode)

	if currentSize < cache.maxSize {
		var newNode = &DoublyLinkedListNode{
			key:  key,
			val:  value,
			prev: cache.tail,
			next: nil,
		}
		cache.tail.next = newNode
		cache.tail = newNode
		cache.keyToListNode[key] = newNode
		return
	}

	// Case where cache is full and requires eviction
	var newNode = &DoublyLinkedListNode{
		key:  key,
		val:  value,
		prev: cache.tail,
		next: nil,
	}
	cache.tail.next = newNode
	cache.tail = newNode
	delete(cache.keyToListNode, cache.head.key)
	cache.head = cache.head.next
	cache.head.prev = nil
	cache.keyToListNode[key] = newNode
}

// The second return value indicates whether or not the key was found
// in the cache.
func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	var node, keyExists = cache.keyToListNode[key]
	if !keyExists {
		return -1, false
	}
	var val = node.val

	if node.next == nil {
		// It is the tail node
		return val, true
	}

	if node.prev != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	} else {
		// It is the head node
		cache.head = node.next
		cache.head.prev = nil
	}

	cache.tail.next = node
	node.prev = cache.tail
	cache.tail = node

	return val, true
}

// The second return value is false if the cache is empty.
func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	if cache.head == nil {
		return "", false
	}
	return cache.tail.key, true
}

func main() {
	var lruCache = NewLRUCache(3)
	lruCache.InsertKeyValuePair("b", 2)
	lruCache.InsertKeyValuePair("a", 1)
	lruCache.InsertKeyValuePair("c", 3)

	var key1, exists1 = lruCache.GetMostRecentKey()
	log.Println("key:", key1, "exists:", exists1)

	var val2, exists2 = lruCache.GetValueFromKey("a")
	log.Println("val:", val2, "exists:", exists2)

	var key3, exists3 = lruCache.GetMostRecentKey()
	log.Println("key:", key3, "exists:", exists3)

	lruCache.InsertKeyValuePair("d", 4)

	var val4, exists4 = lruCache.GetValueFromKey("b")
	log.Println("val:", val4, "exists:", exists4)

	lruCache.InsertKeyValuePair("a", 5)

	var val5, exists5 = lruCache.GetValueFromKey("a")
	log.Println("val:", val5, "exists:", exists5)

}
