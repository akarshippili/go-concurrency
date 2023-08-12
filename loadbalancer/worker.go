package loadbalancer

type Worker struct {
	Requests    chan int
	NumRequests int
}
