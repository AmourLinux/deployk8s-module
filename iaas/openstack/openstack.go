package openstack

import (
	"github.com/amourlinux/deployk8s-module/iaas"
	"time"
)

type Openstack struct {
	addr, token string
}

func New(addr string) *Openstack {
	return &Openstack{addr: addr}
}

func (o *Openstack) CreateLB() (iaas.LB, error) {
	return &LB{
		Addr: "127.0.0.1",
		port: "6443",
	}, nil
}

func (o *Openstack) CreateServer() (iaas.Server, error) {
	return &Server{
		addr:  "127.0.0.1",
		image: "centos7-docker18u9-kernel5",
	}, nil
}

func (o *Openstack) WaitForSync() {
	time.Sleep(time.Second * 3)
}
