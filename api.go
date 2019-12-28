package main

import (
	"flag"
	"github.com/amourlinux/deployk8s-module/cluster"
	"k8s.io/klog"
	"net/http"
)

func init() {
	klog.InitFlags(nil)
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "4")
	flag.Parse()
}

func main() {
	go func() {
		if err := http.ListenAndServe(":80", nil); err != nil {
			panic(err)
		}
	}()

	o := &cluster.Options{
		Name:     "dev",
		IaasType: "openstack",
	}

	c := cluster.New(o)
	err := c.Init()
	if err != nil {
		panic(err)
	}

	//return c
}
