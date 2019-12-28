package openstack

//encapsulation all
type LB struct {
	Addr string
	port string
}

func (lb *LB) GetLBAddress() string {
	return lb.Addr
}

func (lb *LB) GetPort() string {
	return lb.port
}