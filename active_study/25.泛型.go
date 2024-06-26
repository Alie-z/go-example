package heaps

import (
	"container/heap"
	"fmt"
)

// compare function
type CMP[E any] func(E, E) int

// heapST to implments the interface of "heap.Interface"
type heapST[E any] struct {
	data []E
	cmp  CMP[E]
}

// implments the methods for "heap.Interface"
func (h *heapST[E]) Len() int { return len(h.data) }
func (h *heapST[E]) Less(i, j int) bool {
	v := h.cmp(h.data[i], h.data[j])
	return v < 0
}
func (h *heapST[E]) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *heapST[E]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	v := append(h.data, x.(E))
	h.data = v
}
func (h *heapST[E]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

// Heap base on generics to build a heap tree for any type
type Heap[E any] struct {
	data *heapST[E]
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[E]) Push(v E) {
	heap.Push(h.data, v)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *Heap[E]) Pop() E {
	return heap.Pop(h.data).(E)
}

func (h *Heap[E]) Element(index int) (e E, err error) {
	if index < 0 || index >= h.data.Len() {
		return e, fmt.Errorf("out of index")
	}
	return h.data.data[index], nil
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[E]) Remove(index int) E {
	return heap.Remove(h.data, index).(E)
}

// NewHeap return Heap pointer and init the heap tree
func NewHeap[E any](t []E, cmp CMP[E]) *Heap[E] {
	ret := heapST[E]{data: t, cmp: cmp}
	heap.Init(&ret)
	return &Heap[E]{&ret}
}
