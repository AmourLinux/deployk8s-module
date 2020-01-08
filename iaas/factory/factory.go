package factory

import (
	"github.com/amourlinux/deployk8s-module/iaas"
	"github.com/amourlinux/deployk8s-module/iaas/openstack"
)

func New(iaasType string) iaas.Iaas {
	return openstack.New("192.168.80.12")
}

/*
var (
	ostack, al, aw iaas.Iaas
)

var (
	ostackOnce = sync.Once{}
	awsOnce    = sync.Once{}
	aliOnce    = sync.Once{}
)

func NewIaas(iaasType iaas.Type) iaas.Iaas {
	ctx, _ := context.WithCancel(context.Background())

	switch iaasType {
	case iaas.OPENSATCK:
		onceBody := func() {
			ostack = openstack.NewOpenstack("127.0.0.1")
			go ostack.Run(ctx)
			ostack.WaitForSync(ctx)
		}

		ostackOnce.Do(onceBody)
		return ostack

	case iaas.ALIYUN:
	case iaas.AWS:
	}

	return nil
}
*/
