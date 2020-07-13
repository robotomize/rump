package server

type HandleFn func([]byte) error

type Server interface {
	ServeUDP() error
	ServeGRPC() error
}
