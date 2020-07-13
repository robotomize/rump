package client

import (
	"context"
	"rump/internal/logging"
	"rump/internal/srvenv"
	"rump/internal/utils"
	"syscall"
)

type UDPClient struct {
	env      *srvenv.SrvEnv
	sockAddr syscall.Sockaddr
	ctx      context.Context
	logger   logging.Logger
	fd       int
}

func NewUDPClient(ctx context.Context, strAddr string) (*UDPClient, error) {
	addr, err := utils.ParseAddr(strAddr)
	if err != nil {
		return nil, err
	}
	return &UDPClient{sockAddr: addr, logger: logging.FromContext(ctx)}, nil
}

func (u *UDPClient) Open() error {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return err
	}

	if err := syscall.Connect(fd, u.sockAddr); err != nil {
		return err
	}
	u.fd = fd
	return nil
}

func (u *UDPClient) Write(bytes []byte) error {
	_, err := syscall.Write(u.fd, bytes)
	if err != nil {
		return err
	}
	return nil
}

func (u *UDPClient) Close() error {
	if err := syscall.Close(u.fd); err != nil {
		return err
	}
	return nil
}
