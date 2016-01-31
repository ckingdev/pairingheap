package minpheap

import (
	"fmt"
)

type MinPairingHeap struct {
	head     *node
	contains map[interface{}]*node
	size     int
}

type node struct {
	val                    interface{}
	key                    float32
	child, sibling, parent *node
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
		m1child := m1.child
		m1.child = m2
		m2.parent = m1
		m2.sibling = m1child
		return m1
	}
	m2child := m2.child
	m2.child = m1
	m1.parent = m2
	m1.sibling = m2child
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
	if m.head != nil {
		m.head.parent = nil
	}
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

func (m *MinPairingHeap) DecreaseKey(val interface{}, newKey float32) error {
	node, ok := m.contains[val]
	if !ok {
		return fmt.Errorf("Could not find node to update.")
	}
	detach(node)
	node.key = newKey
	m.head = mergeNodes(m.head, node)
	return nil
}

func (m *MinPairingHeap) PeekAtVal(val interface{}) (float32, bool) {
	node, ok := m.contains[val]
	if !ok {
		return 0, false
	}
	return node.key, true
}

func detach(n *node) {
	if n.parent == nil {
		return
	}
	iter := n.parent.child
	if iter == n {
		n.parent.child = n.sibling
		n.parent = nil
		return
	}
	for iter != nil && iter.sibling != n {
		iter = iter.sibling
	}
	iter.sibling = n.sibling
	n.parent = nil
}
