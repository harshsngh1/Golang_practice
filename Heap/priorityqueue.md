# PriorityQueue in Golang

Heap ek special tree-based data structure hai jo priority queue ko efficiently manage karne mein use hota hai. Priority queue mein elements ko priority ke hisaab se manage kiya jata hai.
- Min-Heap: Sabse chhota element root par hota hai.
- Max-Heap: Sabse bada element root par hota hai.

### `container/heap` Package ka Use
Go mein heap implement karne ke liye container/heap package use hota hai. Yeh package aapko min-heap aur max-heap dono implement karne ki flexibility deta hai.  
Heap ko use karne ke liye aapko heap.Interface ko implement karna padta hai. Is interface ke methods hain:
1. Len(): Heap mein elements ki count return karta hai.
2. Less(i, j int) bool: Yeh method define karta hai ki kaunsa element priority ke hisaab se pehle aayega.
3. Swap(i, j int): Index i aur j wale elements ko swap karta hai.
4. Push(x interface{}): Naya element heap mein add karta hai.
5. Pop() interface{}: Heap se element ko remove karta hai aur return karta hai.

### Example: Min-Heap
```
package main

import (
	"container/heap"
	"fmt"
)

// IntHeap ek min-heap ko represent karta hai
type IntHeap []int

// Heap interface ko implement karna
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap condition
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push method ko implement karna
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop method ko implement karna
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5} // Initial heap
	heap.Init(h)           // Heap ko initialize karna
	heap.Push(h, 3)        // Naya element add karna
	fmt.Printf("Minimum: %d\n", (*h)[0]) // Root element print karna
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // Elements ko pop karna
	}
}
```
- IntHeap: Yeh ek slice hai jo heap ko represent karta hai.
- Len, Less, Swap: Yeh methods heap interface ko fulfill karte hain.
- Push, Pop: Naye elements ko add aur existing elements ko remove karne ke methods hain.
- heap.Init: Heap ko initialize karta hai.
- heap.Push: Element ko heap mein add karta hai.
- heap.Pop: Element ko heap se remove aur return karta hai.

Max-Heap Banane ke Liye  
```
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-Heap condition
```

### Doing with custom data structure
```
type DriverHeap []*CabDriver

func (h DriverHeap) Len() int {
	return len(h)
}
func (h DriverHeap) Less(i, j int) bool {
	return h[i].AvgRating > h[j].AvgRating
}
func (h DriverHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *DriverHeap) Push(x interface{}) {
	*h = append(*h, x.(*CabDriver))
}
func (h *DriverHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
```