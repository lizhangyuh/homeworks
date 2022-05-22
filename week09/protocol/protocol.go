package protocol

import (
	"encoding/binary"
	"errors"
	"main/week09/pkg/bufio"
)

const (
	PackSize      = 4
	HeaderSize    = 2
	VerSize       = 2
	OpSize        = 4
	SeqSize       = 4
	RawHeaderSize = PackSize + HeaderSize + VerSize + OpSize + SeqSize

	HeaderOffset = 0 + PackSize
	VerOffset    = HeaderOffset + HeaderSize
	OpOffset     = VerOffset + VerSize
	SeqOffset    = OpOffset + OpSize
	BodyOffset   = SeqOffset + SeqSize
)

type Pack struct {
	Version uint16
	Option  uint32
	Seq     uint32
	Body    []byte
}

func Read(pr *bufio.Reader) (*Pack, error) {
	buf, err := pr.Pop(RawHeaderSize)
	if err != nil {
		return nil, err
	}

	packLen := binary.BigEndian.Uint32(buf[0:HeaderOffset])
	headerLen := binary.BigEndian.Uint16(buf[HeaderOffset:VerOffset])
	version := binary.BigEndian.Uint16(buf[VerOffset:OpOffset])
	option := binary.BigEndian.Uint32(buf[OpOffset:SeqOffset])
	seq := binary.BigEndian.Uint32(buf[SeqOffset:BodyOffset])

	if RawHeaderSize != headerLen {
		return nil, errors.New("header length is wrong")
	}

	bodyLen := int32(packLen) - int32(headerLen)
	if bodyLen < 0 {
		return nil, errors.New("body length is wrong")
	} else {
		body, err := pr.Pop(int(bodyLen))

		if err != nil {
			return nil, err
		}

		return &Pack{
			Version: version,
			Seq:     seq,
			Option:  option,
			Body:    body,
		}, nil
	}

}
