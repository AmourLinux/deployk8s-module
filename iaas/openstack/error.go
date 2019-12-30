package openstack

import "golang.org/x/xerrors"

var (
	ErrHttpConnect = xerrors.New("failed to connect iaas")
)
