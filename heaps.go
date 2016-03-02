// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	fmt.Println("initial heap", h)
	heap.Init(h)
	fmt.Println("initalized heap", h)
	heap.Push(h, 3)
	fmt.Println("push 3 to heap:", h)
	heap.Push(h, 10)
	fmt.Println("push 10 to heap:", h)

	// a := heap.Pop(h)
	// fmt.Println(a)
	// fmt.Println(reflect.TypeOf(a))
	// fmt.Println(a + 5)
	// a = heap.Pop(h).(int)
	// fmt.Println(a)
	// a = heap.Pop(h).(int)
	// fmt.Println(a)

	newThing := 5
	for h.Len() > 0 {
		a := heap.Pop(h).(int)
		fmt.Println("what is a?", a)
		if a >= newThing {
			heap.Push(h, a)
			break
		}
		fmt.Println(h)
	}
	fmt.Println("final heap:", h)

	// Alt that isn't super readable
	for a := heap.Pop(h).(int); a < 5; a = heap.Pop(h).(int) {
		fmt.Println(a)
	}

}
