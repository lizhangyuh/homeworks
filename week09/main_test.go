package main

import (
	"encoding/binary"
	"fmt"
	"main/week09/pkg/bufio"
	"main/week09/protocol"
	"net"
	"testing"
)

func TestMain(t *testing.T) {
	//  test
	c, err := net.Dial("tcp", "127.0.0.1:10000")

	if err != nil {
		t.Errorf("dial error: %v\n", err)
	}

	body := []byte("hello")
	packLen := protocol.RawHeaderSize + len(body)

	wr := bufio.NewWriter(c)

	buf := make([]byte, protocol.BodyOffset)

	binary.BigEndian.PutUint32(buf[0:], uint32(packLen))
	binary.BigEndian.PutUint16(buf[protocol.HeaderOffset:], uint16(protocol.RawHeaderSize))
	binary.BigEndian.PutUint16(buf[protocol.VerOffset:], 1)
	binary.BigEndian.PutUint32(buf[protocol.OpOffset:], 2)
	binary.BigEndian.PutUint32(buf[protocol.SeqOffset:], 3)

	wr.Write(buf)
	fmt.Printf("input body: %v\n", body)
	wr.Write(body)
	wr.Flush()

	fmt.Println("called")
}
