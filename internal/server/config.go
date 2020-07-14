package server

type Config struct {
	UDPAddr  string `env:"UDP_ADDR, default=:5555"`
	GRPCAddr string `env:"GRPC_ADDR, default=:5577"`
}

func (c *Config) GetUDPAddr() string {
	return c.UDPAddr
}

func (c *Config) GetGRPCAddr() string {
	return c.GRPCAddr
}
