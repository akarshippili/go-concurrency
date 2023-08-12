package loadbalancer

type pool []*Worker

type Balancer struct {
	Pool pool
	done chan *Worker
}

func (balance *Balancer) Balance() {
	for {
	}
}
func (balance *Balancer) Assign() {}
func (balance *Balancer) Done()   {}
