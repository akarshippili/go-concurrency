package main

import (
	"fmt"

	"github.com/akarshippili/go-concurrency/loadbalancer"
)

func main() {
	balancer := loadbalancer.Listen(10)
	out := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			balancer.Queue <- loadbalancer.Request{
				Fn:  fib,
				Arg: 10,
				C:   out,
			}
		}
	}()

	for {
		res := <-out
		fmt.Println("response ", res)
	}
}

func fib(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return fib(num-1) + fib(num-2)
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
