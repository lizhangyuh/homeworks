package main

import (
	"fmt"
	"log"
	"main/week09/pkg/bufio"
	"main/week09/protocol"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	fmt.Printf("server starting: %v\n", listener.Addr())
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("accept error: %v\n", err)
			continue
		}

		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	rd := bufio.NewReader(conn)
	// wr := bufio.NewWriter(conn)

	pack, err := protocol.Read(rd)
	if err != nil {
		fmt.Printf("read error: %v\n", err)
	}

	fmt.Printf("pack: %v\n", pack)
	fmt.Printf("body: %v\n", string(pack.Body))
}
