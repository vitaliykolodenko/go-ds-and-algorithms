package main

//Heap store elements and it's length for heap ds
type Heap struct {
	data   []int
	length int
}

// NewHeap creates new Heap with predefined capacity
func NewHeap(size int) *Heap {
	h := Heap{
		data:   make([]int, 0, size),
		length: 0,
	}
	return &h
}

//Add adds and item to the Heap
func (h *Heap) Add(item int) {
	h.data = append(h.data, item)
	h.length++
	i := h.length - 1
	for i > 0 && h.data[i] > h.data[i/2] {
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		i /= 2
	}
}

//Max gets current max in the Heap
func (h *Heap) Max() int {
	return h.data[0]
}

//ExtractMax returns current max and removes it from the Heap
func (h *Heap) ExtractMax() int {
	max := h.data[0]
	h.data[0], h.data[h.length-1] = h.data[h.length-1], h.data[0]
	h.length--
	heapify(h, 0)
	return max
}

func heapify(h *Heap, n int) {
	l := (n+1)*2 - 1
	r := (n + 1) * 2

	if l < h.length && h.data[n] < h.data[l] {
		h.data[n], h.data[l] = h.data[l], h.data[n]
		heapify(h, l)
	} else if r < h.length && h.data[n] < h.data[r] {
		h.data[n], h.data[r] = h.data[r], h.data[n]
		heapify(h, r)
	}
}
