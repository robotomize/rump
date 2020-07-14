package msgpack

import (
	"github.com/ugorji/go/codec"
	"sync"
)

func New() *MsgPack {
	m := &MsgPack{
		encBuf: make([]byte, 1500),
		decBuf: make([]byte, 1500),
	}
	m.enc = codec.NewEncoderBytes(&m.encBuf, &m.mh)
	m.dec = codec.NewDecoderBytes(m.decBuf, &m.mh)
	return m
}

type MsgPack struct {
	mu     sync.RWMutex
	encBuf []byte
	decBuf []byte
	enc    *codec.Encoder
	dec    *codec.Decoder
	mh     codec.MsgpackHandle
}

func (m *MsgPack) Encode(v interface{}) ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.encBuf = m.encBuf[:0]
	m.enc.ResetBytes(&m.encBuf)
	if err := m.enc.Encode(v); err != nil {
		return nil, err
	}
	b := make([]byte, len(m.encBuf))
	copy(b, m.encBuf)
	return b, nil
}

func (m *MsgPack) Decode(bytes []byte, v interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.dec.ResetBytes(bytes)
	if err := m.dec.Decode(v); err != nil {
		return err
	}
	m.decBuf = m.decBuf[:0]
	return nil
}
