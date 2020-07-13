package server

type Config struct {
	SyncAddr string `env:"SYNC_ADDR, default=:5555"`
	RcvAddr  string `env:"RCV_ADDR, default=:5577"`
}
