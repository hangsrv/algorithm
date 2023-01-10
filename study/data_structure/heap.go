package data_structure

import (
	"container/heap"
	"fmt"
)

func InitHeap() {
	h := IntHeap{1, 3, 5, 2, 4}
	heap.Init(&h)

	for len(h) > 0 {
		fmt.Println(heap.Pop(&h))
	}
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	l := len(old)
	pop := old[l-1]
	*h = old[0 : l-1]
	return pop
}
