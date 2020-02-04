package design

type Node struct {
	Key  int
	Val  int
	Cnt  int
	Prev *Node
	Next *Node
}

type DLinkedList struct {
	head *Node
	tail *Node
	Size int
}

func CreateDL() *DLinkedList {
	head, tail := &Node{}, &Node{}
	head.Next, tail.Prev = tail, head
	return &DLinkedList{
		head: head,
		tail: tail,
		Size: 0,
	}
}

func (this *DLinkedList) remove(n *Node) {
	dummy := n.Prev
	dummy.Next = n.Next
	n.Next.Prev = dummy
	this.Size -= 1
}

func (this *DLinkedList) insert(node *Node) {
	node.Prev = this.tail.Prev
	this.tail.Prev.Next = node
	node.Next = this.tail
	this.tail.Prev = node
	this.Size += 1
}

func (this *DLinkedList) pop() *Node {
	v := this.head.Next
	this.remove(this.head.Next)
	return v
}

type LFUCache struct {
	min      int
	capacity int
	hash     map[int]*Node
	freqHash map[int]*DLinkedList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{-1, capacity, make(map[int]*Node), make(map[int]*DLinkedList)}
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.hash[key]
	if ok {
		// exist key, update frequency
		freq := v.Cnt

		this.freqHash[freq].remove(v)
		if this.freqHash[freq+1] == nil {
			this.freqHash[freq+1] = CreateDL()
		}
		this.freqHash[freq+1].insert(v)
		if freq == this.min && this.freqHash[freq].Size == 0 {
			this.min += 1
		}
		v.Cnt += 1
		return v.Val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}
	v, ok := this.hash[key]
	if ok {
		// key exist, update value and freq
		freq := v.Cnt
		this.freqHash[freq].remove(v)
		if this.freqHash[freq+1] == nil {
			this.freqHash[freq+1] = CreateDL()
		}
		this.freqHash[freq+1].insert(v)
		if freq == this.min && this.freqHash[freq].Size == 0 {
			this.min += 1
		}
		v.Val = value
		v.Cnt += 1
	} else {
		// key not exist
		if len(this.hash) >= this.capacity {
			tmp := this.freqHash[this.min].pop()
			delete(this.hash, tmp.Key)
		}
		this.min = 1
		node := &Node{key, value, 1, nil, nil}
		if this.freqHash[1] == nil {
			this.freqHash[1] = CreateDL()
		}
		this.hash[key] = node
		this.freqHash[1].insert(node)
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
