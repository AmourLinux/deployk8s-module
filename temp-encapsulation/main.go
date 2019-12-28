package main

import (
	"fmt"
	"github.com/amourlinux/deployk8s-module/iaas/openstack"
	"reflect"
)

func main() {
	lb := &openstack.LB{
		Addr: "0.0.0.0",
	}
	fmt.Printf("%+v\n", lb)

	v := reflect.ValueOf(lb)
	fmt.Println(v.CanSet())
	fmt.Println(v.Elem().CanSet())

	ve := v.Elem()
	fmt.Println(ve.CanSet())
	fmt.Println(ve.Kind(), ve.Type().Kind(), ve.Type())

	fmt.Println(ve.NumField())
	//fmt.Println(v.NumField())
	fmt.Println(ve.Field(0), ve.Field(1))
	if ve0 := ve.Field(0); ve0.CanSet() {
		fmt.Println(ve0.Kind())
		ve0.SetString("127.0.0.1")
	}
	fmt.Println(ve.Field(0))
	fmt.Printf("%+v\n", lb)

	fmt.Println(ve.Field(1).CanSet())
}
