package merr

import (
	"fmt"
	"golang.org/x/xerrors"
	"k8s.io/klog"
)

var (
	ErrConnectIaas = xerrors.New("failed to connect iaas")
	ErrReadBody    = xerrors.New("failed to read resp body")
	ErrHttpResp    = xerrors.New("invalid http response code")

	ErrMakeRequest   = xerrors.New("failed to make request")
	ErrReqConflict   = xerrors.New("request conflict")
	ErrInvalidUrl    = xerrors.New("invalid url")
	ErrInvalidString = xerrors.New("invalid string")

	ErrJsonMarshal   = xerrors.New("failed to json marshal")
	ErrJsonUnmarshal = xerrors.New("failed to json unmarshal")

	ErrInvalidIaasType = xerrors.New("invalid iaas type")
	ErrDBConnect       = xerrors.New("failed to connect DB")
	ErrInvalidateRole  = xerrors.New("invalid role for server")
)

func main() {
	//xerrors.Formatter()
	//xerrors.FormatError()
	e1 := xerrors.Errorf("http error: %w", ErrConnectIaas)
	klog.Errorf("%v\n", e1)
	fmt.Printf("%+v\n", e1)
	eo := xerrors.Opaque(e1)
	fmt.Println(eo)

	fmt.Println()
	m1, ok := e1.(xerrors.Wrapper)
	if !ok {
		fmt.Println("Not imp")
	} else {
		fmt.Println(m1)
	}

	ok = xerrors.Is(e1, ErrConnectIaas)
	if !ok {
		fmt.Println("Isn't")
	} else {
		fmt.Println("Is")
	}

	ok = xerrors.As(e1, &ErrConnectIaas)
	if !ok {
		fmt.Println("Asn't")
	} else {
		fmt.Println("As")
	}
}
