package loadbalancer

type Request struct {
	Fn func() int
	C  chan int
}
