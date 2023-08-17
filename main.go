package main

import (
	"fmt"
	"time"

	"github.com/akarshippili/go-concurrency/loadbalancer"
)

func main() {
	balancer := loadbalancer.Listen(10)
	out := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			balancer.Queue <- loadbalancer.Request{
				Fn:  ioWork,
				Arg: 30,
				C:   out,
			}
		}
	}()

	for {
		res := <-out
		fmt.Println("response ", res)
	}
}

// io bound task
func ioWork(x int) int {
	time.Sleep(time.Millisecond * 2)
	return 0
}

// cpu bound task
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
