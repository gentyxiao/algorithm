package exercise

type DListNode struct {
	Pre        *DListNode
	Next       *DListNode
	Key, Value int
}

type LRUCache struct {
	head     *DListNode
	tail     *DListNode
	hashMap  map[int]*DListNode
	capacity int
	size     int
}

func Constructor(capacity int) LRUCache {
	head := &DListNode{
		Pre:  nil,
		Next: nil,
	}
	tail := &DListNode{
		Pre:  nil,
		Next: nil,
	}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		hashMap:  make(map[int]*DListNode),
		capacity: capacity,
		size:     0,
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	got, ok := l.hashMap[key]
	if ok {
		// 元素存在，需要将元素插入到链表的头部
		// 删除当前结点
		got.Pre.Next = got.Next
		got.Next.Pre = got.Pre
		// 移动到头部
		got.Next = l.head.Next
		got.Pre = l.head
		l.head.Next.Pre = got
		l.head.Next = got
		return got.Value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	// 1 检查元素是否存在
	got, ok := l.hashMap[key]
	if !ok {
		// 元素不存在，直接插入到头部
		item := &DListNode{
			Value: value,
			Key:   key,
		}
		item.Next = l.head.Next
		item.Pre = l.head
		l.head.Next.Pre = item
		l.head.Next = item
		l.size += 1
		l.hashMap[key] = item
		// 元素数量超出最大容量
		if l.size > l.capacity {
			// 删除链表末端元素
			node := l.tail.Pre
			node.Pre.Next = l.tail
			node.Next.Pre = node.Pre
			delete(l.hashMap, node.Key)
			l.size--
		}
	} else {
		// 元素存在，需要将元素插入到链表的头部
		// 删除当前结点
		got.Key = key
		got.Value = value

		got.Pre.Next = got.Next
		got.Next.Pre = got.Pre
		// 移动到头部
		got.Next = l.head.Next
		got.Pre = l.head
		l.head.Next.Pre = got
		l.head.Next = got
	}
}
