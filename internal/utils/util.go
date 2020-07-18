package utils

import (
	"bytes"
	"net"
	"sync"
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

var bytesBuffer = sync.Pool{
	New: func() interface{} { return &bytes.Buffer{} },
}

func GetBuffer() (p *bytes.Buffer) {
	ifc := bytesBuffer.Get()
	if ifc != nil {
		p = ifc.(*bytes.Buffer)
	}
	return
}

func PutBuffer(p *bytes.Buffer) {
	bytesBuffer.Put(p)
}
