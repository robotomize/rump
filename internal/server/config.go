package server

type Kind uint8

const (
	KindUDP Kind = iota
	KindTCP
)

type Config struct {
	Addr string `env:"GAME_STATER_ADDR, default=:5555"`
	Kind Kind   `env:"GAME_STATER_KIND, default=0"`
}

func (c *Config) IsUDP() bool {
	return c.Kind == KindUDP
}

func (c *Config) IsTCP() bool {
	return c.Kind == KindTCP
}
