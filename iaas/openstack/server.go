package openstack

type Server struct {
	addr 	string
	image 	string
}

func (s *Server) GetServerAddress() string {
	return s.addr
}