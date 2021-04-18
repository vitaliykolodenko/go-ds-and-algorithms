package main

import "fmt"

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
	i := h.length
	for i > 1 && h.data[i-1] > h.data[i/2-1] {
		h.data[i-1], h.data[i/2-1] = h.data[i/2-1], h.data[i-1]
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
	fmt.Printf("Heap state after extract: length - %d, heap - %v \n", h.length, h.data)
	return max
}

//Sort sorts slice d using heap sort algorithm
func Sort(d []int) []int {
	res := make([]int, len(d))
	h := NewHeap(len(res))

	for _, v := range d {
		h.Add(v)
	}

	fmt.Printf("Heap after insert: %v\n", h.data)

	for i := len(d) - 1; i >= 0; i-- {
		res[i] = h.ExtractMax()
		fmt.Printf("Adding item to result at %d, val: %d \n", i, res[i])
	}
	fmt.Printf("Result is: %v\n", res)
	return res
}

func heapify(h *Heap, n int) {
	l := (n+1)*2 - 1
	r := (n + 1) * 2

	if r < h.length {
		if h.data[l] > h.data[r] && h.data[l] > h.data[n] {
			h.data[n], h.data[l] = h.data[l], h.data[n]
			heapify(h, l)
		} else if h.data[r] > h.data[n] {
			h.data[n], h.data[r] = h.data[r], h.data[n]
			heapify(h, r)
		}
	} else {
		if l < h.length && h.data[n] < h.data[l] {
			h.data[n], h.data[l] = h.data[l], h.data[n]
			heapify(h, l)
		}
	}
}
