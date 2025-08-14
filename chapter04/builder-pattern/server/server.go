package server

import (
	"log"
	"time"
)

type Server struct {
	param serverParam
}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewBuilder(host string, port int) *serverParam {
	return &serverParam{host: host, port: port}
}

func (sb *serverParam) Timeout(timeout time.Duration) *serverParam {
	sb.timeout = timeout
	return sb
}

func (sb *serverParam) Logger(logger *log.Logger) *serverParam {
	sb.logger = logger
	return sb
}

func (sb *serverParam) Build() *Server {
	svr := &Server{
		param: *sb,
	}
	return svr
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Println("server started")
	}
	// do something
	return nil
}
