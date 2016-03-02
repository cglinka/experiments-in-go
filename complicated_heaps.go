package main

import (
	"container/heap"
	"fmt"
)

type blob struct {
	version int
}

type BlobHeap []blob

func (h BlobHeap) Len() int           { return len(h) }
func (h BlobHeap) Less(i, j int) bool { return h[i].version < h[j].version }
func (h BlobHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BlobHeap) Push(x interface{}) {
	*h = append(*h, x.(blob))
}

func (h *BlobHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &BlobHeap{blob{6}, blob{9}, blob{1}, blob{8}, blob{3}}
	heap.Init(h)
	fmt.Println(h)
	heap.Push(h, blob{10})
	fmt.Println(h)
	heap.Push(h, blob{2})
	fmt.Println(h)

	for h.Len() > 0 {
		a := heap.Pop(h).(blob)
		fmt.Println("What is a?", a.version)
	}
}
