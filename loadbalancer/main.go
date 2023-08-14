package loadbalancer

import (
	"github.com/akarshippili/go-concurrency/heap"
)

func Listen(numWorkers int) *Balancer {

	lb := Balancer{
		Queue:  make(chan Request),
		DoneCh: make(chan *Worker),
		Pool:   heap.GetHeap[*Worker](),
	}

	for num := 0; num < numWorkers; num++ {
		worker := GetWorker()
		lb.Pool.Add(worker)
		go worker.Work(&lb)
	}

	go func() { lb.Balance() }()
	return &lb
}
