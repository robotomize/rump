package srvenv

type ServerConfig interface {
	GetUDPAddr() string
	GetGRPCAddr() string
}

type SrvEnv struct {
	SrvConfig ServerConfig
}
