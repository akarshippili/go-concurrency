package main

import (
	"fmt"
	"strconv"

	"github.com/akarshippili/go-concurrency/heap"
)

func main() {

	pq := heap.GetHeap[*heap.HeapString]()

	for i := 9; i >= 0; i -= 1 {
		pq.Add(heap.GetHeapString(strconv.Itoa(i)))
		val, err := pq.Peek()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("heap peek : %v \n", *val)
		}
	}

	pq.IncreaseKeyWith(0, heap.GetHeapString("99"))
	pq.Delete(0)

	for !pq.IsEmpty() {
		val, err := pq.Pop()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("heap pop : %v \n", *val)
		}
	}
}

// func main() {
// 	heap := heap.GetHeap[HeapString]()
// }
