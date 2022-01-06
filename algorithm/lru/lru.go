package lru

type linkNode struct {
	pre  *linkNode
	next *linkNode
	key  int
	val  int
}

type cache struct {
	m    map[int]*linkNode //指向哈希表的指针
	cap  int               //长度
	head *linkNode         //两个哨兵
	tail *linkNode
}

func New(capacity int) cache {
	head := &linkNode{
		pre:  nil,
		next: nil,
		key:  0,
		val:  0,
	}

	tail := &linkNode{
		pre:  nil,
		next: nil,
		key:  0,
		val:  0,
	}

	head.next = tail
	tail.pre = head
	return cache{
		m:    make(map[int]*linkNode),
		cap:  capacity,
		head: head,
		tail: tail,
	}
}

func (l *cache) Get(key int) int {
	cache := l.m
	if node, ok := cache[key]; ok {
		l.moveToHead(node)
		return node.val
	}

	return -1
}

func (l *cache) Put(key int, val int) {
	if node, ok := l.m[key]; ok {
		node.val = val
		l.moveToHead(node)
		return
	}

	node := new(linkNode)
	node.key = key
	node.val = val

	if len(l.m) == l.cap {
		delete(l.m, l.tail.pre.key)
		l.popToTail()
	}
	l.addNode(node)
}

func (l *cache) moveToHead(node *linkNode) {
	l.removeNode(node)
	l.addNode(node)
}

func (l *cache) popToTail() {
	l.removeNode(l.tail.pre)
}

// addNode 在头部添加节点
func (l *cache) addNode(node *linkNode) {
	l.m[node.key] = node

	node.pre = l.head
	node.next = l.head.next

	l.head.next.pre = node
	l.head.next = node
}

// removeNode 移除尾部节点
func (l *cache) removeNode(node *linkNode) {
	delete(l.m, node.key)

	pre := node.pre
	next := node.next

	pre.next = next
	next.pre = pre
}
