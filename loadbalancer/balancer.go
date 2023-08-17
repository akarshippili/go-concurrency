package loadbalancer

import (
	"fmt"
	"time"

	"github.com/akarshippili/go-concurrency/heap"
)

type Balancer struct {
	Pool   *heap.Heap[*Worker]
	Queue  chan Request
	DoneCh chan *Worker
}

func (balancer *Balancer) Balance() {
	for {
		select {
		case request := <-balancer.Queue:
			balancer.Assign(request)
		case worker := <-balancer.DoneCh:
			balancer.Done(worker)
		default:
		}
	}
}

func (balancer *Balancer) Assign(req Request) {
	worker, err := balancer.Pool.Peek()
	if err != nil {
		fmt.Println("error while assigning task to worker")
	}

	if len((*worker).Requests) == cap((*worker).Requests) {
		go func() { balancer.Queue <- req }()
		time.Sleep(time.Second)
		return
	}

	worker, err = balancer.Pool.Pop()
	if err != nil {
		fmt.Println("error while assigning task to worker")
	}

	(*worker).NumRequests += 1
	(*worker).Requests <- req
	balancer.Pool.Add(*worker)
}

func (balancer *Balancer) Done(worker *Worker) {
	workerPool := balancer.Pool
	worker.NumRequests -= 1
	worker.NumRequestsHandled += 1
	workerPool.DecreaseKeyWith(worker.GetIndex(), worker)
}

func (balancer *Balancer) Log() {
	for _, worker := range balancer.Pool.Arr {
		fmt.Printf(" %v / %v ", worker.GetNumOfReuestsHandled(), worker.NumRequests)
	}
	fmt.Printf("\n")
}
