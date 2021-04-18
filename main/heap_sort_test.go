package main

import "testing"

func TestCreation(t *testing.T) {
	h := NewHeap(3)

	if h.length != 0 {
		t.Fatalf("Newly created Heap is not 0 size, size: %d", h.length)
	}

	if len(h.data) != 0 {
		t.Fatalf("Newly created Heap has some items, len: %d", len(h.data))
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

func TestHeapExtractMax(t *testing.T) {
	h := NewHeap(3)

	h.Add(1)
	h.Add(2)
	h.Add(5)

	max := h.ExtractMax()

	if max != 5 {
		t.Fatalf("Heap did not return max, returned: %d", max)
	}

	max = h.Max()

	if max != 2 {
		t.Fatalf("Heap did not extract previous max, returned: %d", max)
	}

	if h.length != 2 {
		t.Fatalf("Heap did not extract previous max, length still: %d", h.length)
	}
}
