package exercise

type DListNode struct {
	Pre   *DListNode
	Next  *DListNode
	Value int
}

type LRUCache struct {
	Head     *DListNode
	Tail     *DListNode
	HashMap  map[int]*DListNode
	capacity int
	n        int
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
		HashMap:  make(map[int]*DListNode),
		capacity: capacity,
		n:        0,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	got, ok := this.HashMap[key]
	if ok {
		// 元素存在，需要将元素插入到链表的头部
		got.Pre.Next = got.Next
		got.Next.Pre = got.Pre

		got.Next = this.Head.Next
		got.Pre = this.Head
		this.Head.Next.Pre = got
		this.Head.Next = got
		return got.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 1 检查元素是否存在
	got, ok := this.HashMap[key]
	if !ok {
		// 元素不存在，直接插入到头部
		item := &DListNode{
			Value: value,
		}
		item.Next = this.Head.Next
		item.Pre = this.Head
		this.Head.Next.Pre = item
		this.Head.Next = item
		this.n += 1
		// 元素数量超出最大容量
		if this.n > this.capacity {
			// 删除链表末端元素
			this.Tail.Pre.Pre.Next = this.Tail
			this.Tail.Pre = this.Tail.Pre.Pre
		}
	} else {
		// 元素存在，需要将元素插入到链表的头部
		got.Pre.Next = got.Next
		got.Next.Pre = got.Pre

		got.Next = this.Head.Next
		got.Pre = this.Head
		this.Head.Next.Pre = got
		this.Head.Next = got
	}
}
