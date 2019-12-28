package factory

import (
	"github.com/amourlinux/deployk8s-module/iaas"
	"github.com/amourlinux/deployk8s-module/iaas/openstack"
)

func New(iaasType string) iaas.Iaas {
	return openstack.New("192.168.80.12")
}