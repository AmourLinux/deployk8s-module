package iaas

type Iaas interface {
	CreateLB() (LB, error)
	CreateServer() (Server, error)
	//WaitForSync()
	//UpdateLBForAddUpstream()
}

type LB interface {
	GetLBAddress() string
	GetPort() string
}

type Server interface {
	GetServerAddress() string
}

