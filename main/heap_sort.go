package main

import "fmt"

type heap struct {
	data   []int
	length int
}

func NewHeap(size int) *heap {
	h := heap{
		data:   make([]int, 0, size),
		length: 0,
	}
	return &h
}

func (h *heap) Add(item int) {
	h.data = append(h.data, item)
	h.length++
	i := h.length - 1
	for i > 0 && h.data[i] > h.data[i/2] {
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		i /= 2
	}

	fmt.Printf("Heap after add: %v \n", h.data)
}

func (h *heap) Max() int {
	return h.data[0]
}

func (h *heap) ExtractMax() int {
	max := h.data[0]
	h.data[0], h.data[h.length-1] = h.data[h.length-1], h.data[0]
	h.length--
	return max
}
