package minpheap

type MinPairingHeap struct {
	head     *node
	contains map[interface{}]*node
	size     int
}

type node struct {
	val            interface{}
	key            float32
	child, sibling *node
}

func New() *MinPairingHeap {
	return &MinPairingHeap{
		contains: make(map[interface{}]*node),
	}
}

func (m *MinPairingHeap) Peek() (interface{}, float32) {
	if m.head == nil {
		return nil, 0
	}
	return m.head.val, m.head.key
}

func mergeNodes(m1, m2 *node) *node {
	if m1 == nil {
		return m2
	}
	if m2 == nil {
		return m1
	}
	if m1.key < m2.key {
		tmp := m1.child
		m1.child = m2
		m2.sibling = tmp
		return m1
	}
	tmp := m2.child
	m2.child = m1
	m1.sibling = tmp
	return m2
}

func (m *MinPairingHeap) Insert(val interface{}, key float32) {
	tmp := &node{
		val: val,
		key: key,
	}
	m.head = mergeNodes(tmp, m.head)
	m.contains[val] = tmp
	m.size++
}

func (m *MinPairingHeap) Pop() (interface{}, float32) {
	val, key := m.Peek()
	if m.head == nil {
		return nil, 0
	}
	m.head = mergePairs(m.head.child)
	delete(m.contains, val)
	m.size--
	return val, key
}

func mergePairs(n *node) *node {
	if n == nil {
		return nil
	}
	if n.sibling == nil {
		return n
	}
	tmp := n.sibling.sibling
	return mergeNodes(mergeNodes(n, n.sibling), mergePairs(tmp))
}
