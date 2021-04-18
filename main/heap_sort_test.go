package main

import "testing"

func TestCreation(t *testing.T) {
	h := NewHeap(3)

	if h.length != 0 {
		t.Fatalf("Newly created heap is not 0 size, size: %d", h.length)
	}

	if len(h.data) != 0 {
		t.Fatalf("Newly created heap has some items, len: %d", len(h.data))
	}
}

func TestHeapAdd(t *testing.T) {
	h := NewHeap(3)

	h.Add(1)
	h.Add(2)
	h.Add(3)

	if h.length != 3 {
		t.Fatalf("Heap did not increase legnth on addition, size: %d", h.length)
	}
}

func TestHeapMax(t *testing.T) {
	h := NewHeap(3)

	h.Add(1)
	h.Add(2)
	h.Add(5)

	max := h.Max()

	if max != 5 {
		t.Fatalf("Heap did not return max, returned: %d", max)
	}
}
