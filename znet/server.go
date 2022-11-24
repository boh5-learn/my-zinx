package znet

// Server 实现 IServer, Server 服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的 ip 版本
	IPVersion string
	// 服务器监听的 IP
	IP string
	// 服务器监听的端口
	Port int
}

func (s *Server) Start() {
	// TODO implement me
	panic("implement me")
}

func (s *Server) Stop() {
	// TODO implement me
	panic("implement me")
}

func (s *Server) Serve() {
	// TODO implement me
	panic("implement me")
}
