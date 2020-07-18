package encbinary

import (
	"encoding/binary"
	"rump/internal/utils"
)

func New() *EncBinary {
	return &EncBinary{}
}

type EncBinary struct {
}

func (m *EncBinary) Encode(v interface{}) ([]byte, error) {
	b := utils.GetBuffer()
	defer utils.PutBuffer(b)
	defer b.Reset()
	if err := binary.Write(b, binary.LittleEndian, v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (m *EncBinary) Decode(bytes []byte, v interface{}) error {
	b := utils.GetBuffer()
	defer utils.PutBuffer(b)
	defer b.Reset()
	b.Write(bytes)
	if err := binary.Read(b, binary.LittleEndian, v); err != nil {
		return err
	}
	return nil
}
