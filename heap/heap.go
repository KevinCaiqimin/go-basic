package heap

import "fmt"

//HeapData 堆数据
type HeapData interface {
	// Compare 比较两个堆数据，
	// if self > other then return 1
	// if self = other then return 0
	// if self < other then return -1
	Compare(other HeapData) int
}

type HeapNode struct {
	parent *HeapNode
	left   *HeapNode
	right  *HeapNode
	next   *HeapNode //层序遍历下一个
	prev   *HeapNode //层序遍历上一个
	data   HeapData
}

type Heap struct {
	root *HeapNode
	last *HeapNode
	min  bool
}

func NewMinHeap() *Heap {
	heap := &Heap{
		min: true,
	}
	return heap
}

func NewMaxHeap() *Heap {
	heap := &Heap{
		min: false,
	}
	return heap
}

func (h *Heap) LevelPrint() {
	lvRoot := h.root
	for lvRoot != nil {
		p := lvRoot
		for p != nil {
			fmt.Printf("%v\t", p.data)
			p = p.next
		}
		fmt.Printf("\n")
		lvRoot = lvRoot.left
	}
}

func (h *Heap) adjustNodeBottomUp(node *HeapNode) {
	if h.root == nil || h.last == nil {
		return
	}
	if node.parent == nil {
		return
	}
	parent := node.parent
	r := node.data.Compare(parent.data)
	swap := false
	if h.min {
		if r < 0 {
			swap = true
		}
	} else {
		if r > 0 {
			swap = true
		}
	}
	if !swap {
		return
	}
	tmp := parent.data
	parent.data = node.data
	node.data = tmp
	h.adjustNodeBottomUp(parent)
}

func (h *Heap) selectChild(node *HeapNode, smaller bool) *HeapNode {
	if node.left == nil {
		return nil
	}
	if node.right == nil {
		return node.left
	}
	r := node.left.data.Compare(node.right.data)
	if smaller {
		if r < 0 {
			return node.left
		}
		return node.right
	} else {
		if r > 0 {
			return node.left
		}
		return node.right
	}
}

func (h *Heap) adjustNodeTopDown(node *HeapNode) {
	if h.root == nil || h.last == nil {
		return
	}
	if node.left == nil && node.right == nil {
		return
	}
	avaiNode := h.selectChild(node, h.min)
	if avaiNode == nil {
		return
	}

	r := node.data.Compare(avaiNode.data)
	swap := false
	if h.min {
		if r > 0 {
			swap = true
		}
	} else {
		if r < 0 {
			swap = true
		}
	}
	if !swap {
		return
	}
	tmp := avaiNode.data
	avaiNode.data = node.data
	node.data = tmp
	h.adjustNodeTopDown(avaiNode)
}

func (h *Heap) insertNode(node *HeapNode) {
	if h.root == nil {
		h.root = node
		h.last = node
		return
	}
	if h.last == h.root {
		node.parent = h.root
		h.root.left = node
		h.last = node
		return
	}
	if h.last == h.last.parent.left {
		node.parent = h.last.parent
		h.last.parent.right = node
		node.prev = h.last
		h.last.next = node
		h.last = node
		return
	}
	if h.last == h.last.parent.right {
		p := h.last.parent.next
		if p == nil {
			p = h.last
			for p.prev != nil {
				p = p.prev
			}
		}
		node.parent = p
		p.left = node
		h.last = node
		//setup link
		if node.parent.prev != nil {
			node.prev = node.parent.prev.right
			node.parent.prev.right.next = node
		}
		return
	}
}

func (h *Heap) Push(data HeapData) {
	node := &HeapNode{
		data:   data,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	h.insertNode(node)
	h.adjustNodeBottomUp(h.last)
}

func (h *Heap) Get() HeapData {
	if h.root == nil {
		return nil
	}
	return h.root.data
}

func (h *Heap) Pop() HeapData {
	if h.root == nil {
		return nil
	}
	data := h.root.data
	if h.root == h.last {
		h.root = nil
		h.last = nil
	} else {
		h.root.data = h.last.data
		p := h.last.prev
		if p == nil {
			p = h.last.parent
			for p.next != nil {
				p = p.next
			}
		} else {
			p.next = nil
		}
		if h.last == h.last.parent.left {
			h.last.parent.left = nil
		}
		if h.last == h.last.parent.right {
			h.last.parent.right = nil
		}
		h.last = p
		h.adjustNodeTopDown(h.root)
	}
	return data
}
