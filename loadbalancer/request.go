package loadbalancer

type Request struct {
	Fn  func(arg int) int
	C   chan int
	Arg int
}
