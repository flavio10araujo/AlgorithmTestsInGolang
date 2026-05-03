package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	k    int
	heap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)

	kth := KthLargest{
		k:    k,
		heap: h,
	}

	for _, n := range nums {
		kth.Add(n)
	}

	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.heap.Len() < this.k {
		heap.Push(this.heap, val)
		return (*this.heap)[0]
	}

	kthLargest := (*this.heap)[0]

	if val <= kthLargest {
		return kthLargest
	}

	heap.Pop(this.heap)
	heap.Push(this.heap, val)

	return (*this.heap)[0]
}

func main() {
	k := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(k.Add(3))  // returns 4
	fmt.Println(k.Add(5))  // returns 5
	fmt.Println(k.Add(10)) // returns 5
	fmt.Println(k.Add(9))  // returns 8
	fmt.Println(k.Add(4))  // returns 8
}
