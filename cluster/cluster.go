package cluster

import (
	"encoding/json"
	"errors"
	"github.com/amourlinux/deployk8s-module/iaas"
	"github.com/amourlinux/deployk8s-module/iaas/factory"
	"k8s.io/klog"
)

var x iaas.Iaas
var y iaas.Iaas

func init() {
	x = factory.New("openstack")
	x.WaitForSync()
	klog.Infoln("Openstack sync succeed")

	y = factory.New("aws")
	y.WaitForSync()
	klog.Infoln("Aws sync succeed")
}

type Cluster struct {
	Name, Version, IaasType string
}

func New(o *Options) *Cluster {
	c := &Cluster{}
	c.Version = "1.17"

	c.Name = o.Name
	c.IaasType = o.IaasType

	return c
}

type Options struct {
	Name     string
	IaasType string
}

func (c *Cluster) Init() error {
	var i iaas.Iaas
	switch c.IaasType {
	case "openstack":
		i = x
	case "aws":
		i = y
	default:
		return errors.New("invalidate iaas type")
	}

	lb, err := i.CreateLB()
	if err != nil {
		return err
	}
	klog.V(4).Infof("%+v", lb)
	data, _ := json.Marshal(lb)
	json.Unmarshal(data, nil)
	klog.V(4).Infoln(string(data))

	server, err := i.CreateServer()
	if err != nil {
		return err
	}
	klog.V(4).Infof("%+v", server)
	data, _ = json.Marshal(server)
	klog.V(4).Infoln(string(data))

	return nil
}
