package loadbalancer

type Worker struct {
	Requests    chan Request
	NumRequests int
	index       int
}

func GetWorker() *Worker {
	return &Worker{
		Requests:    make(chan Request),
		index:       -1,
		NumRequests: 0,
	}
}

func (worker *Worker) Work(balancer *Balancer) {
	for {
		request := <-worker.Requests
		request.C <- request.Fn(request.Arg)
		balancer.DoneCh <- worker
	}
}

// to complete heapable contract
func abs(i int) int {
	if i >= 0 {
		return i
	}

	return -1 * i
}

func (worker *Worker) Compare(worker2 *Worker) int {
	if worker.NumRequests == worker2.NumRequests {
		return 0
	}

	diff := worker.NumRequests - worker2.NumRequests
	return (diff) / int(abs(diff))
}

func (worker *Worker) SetIndex(index int) {
	worker.index = index
}

func (worker *Worker) GetIndex() int {
	return worker.index
}

func (worker *Worker) GetMin() *Worker {
	return &Worker{
		Requests:    nil,
		NumRequests: 0,
		index:       -1,
	}
}
