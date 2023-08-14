package loadbalancer

import (
	"fmt"

	"github.com/akarshippili/go-concurrency/heap"
)

type Pool *heap.Heap[*Worker]

type Balancer struct {
	Pool   Pool
	Queue  chan Request
	DoneCh chan *Worker
}

func (balancer *Balancer) Balance() {
	for {
		select {
		case request := <-balancer.Queue:
			fmt.Println(request)
		case worker := <-balancer.DoneCh:
			fmt.Println(*worker)
		default:
			fmt.Println("Nothing! Just Hanging Around!!!")
		}
	}
}

func (balancer *Balancer) Assign() {}
func (balancer *Balancer) Done()   {}
