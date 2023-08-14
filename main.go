package main

import (
	"fmt"

	"github.com/akarshippili/go-concurrency/heap"
	"github.com/akarshippili/go-concurrency/loadbalancer"
)

func main() {
	numWorkers := 10
	c := make(chan int)
	var workerPool loadbalancer.Pool = heap.GetHeap[*loadbalancer.Worker]()

	lb := loadbalancer.Balancer{
		Queue:  make(chan loadbalancer.Request),
		DoneCh: make(chan *loadbalancer.Worker),
		Pool:   workerPool,
	}

	fmt.Println(lb, c)

	for num := 0; num < numWorkers; num++ {
		worker := loadbalancer.GetWorker()
		go worker.Work()
	}
}

// ------------------------------------------------------ //
// func main() {

// 	pq := heap.GetHeap[*heap.HeapString]()

// 	for i := 9; i >= 0; i -= 1 {
// 		pq.Add(heap.GetHeapString(strconv.Itoa(i)))
// 		val, err := pq.Peek()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		} else {
// 			fmt.Printf("heap peek : %v \n", *val)
// 		}
// 	}

// 	pq.IncreaseKeyWith(0, heap.GetHeapString("99"))
// 	pq.Delete(0)

// 	for !pq.IsEmpty() {
// 		val, err := pq.Pop()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		} else {
// 			fmt.Printf("heap pop : %v \n", *val)
// 		}
// 	}
// }

// ------------------------------------------------------ //
// func main() {
// 	heap := heap.GetHeap[HeapString]()
// }
