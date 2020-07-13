package utils

import (
	"net"
	"syscall"
)

func ParseAddr(strAddr string) (syscall.Sockaddr, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", strAddr)
	if err != nil {
		return nil, err
	}
	addr := syscall.SockaddrInet4{Port: udpAddr.Port}
	copy(addr.Addr[:], udpAddr.IP.To4())

	return &addr, nil
}
