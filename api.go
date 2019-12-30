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

	//time.Sleep(time.Minute)

	o := &cluster.Options{
		Name:     "dev",
		IaasType: "openstack",
	}

	c := cluster.New(o)
	err := c.Init()
	if err != nil {
		klog.Errorf("Failed to init cluster: %v", err)
		klog.Fatalf("Failed to init cluster: %+v", err)
	}

	//return c
}
