package config

import "os"

const (
	Host EnvVar = "HOST"
	Port EnvVar = "PORT"
)

type Server struct {
	Host string
	Port string
}

func (s *Server) Load() *Server {
	s.Host = os.Getenv(string(Host))
	s.Port = os.Getenv(string(Port))
	return s
}
