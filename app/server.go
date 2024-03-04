package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/lib/parser"
	"github.com/codecrafters-io/redis-starter-go/app/lib/store"
)

const (
	TCP  = "tcp"
	PORT = ":6379"
)

func main() {
	listener, err := net.Listen(TCP, PORT)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	Storage := store.NewStorage()
	fmt.Println("Listening on :6379...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleConnection(conn, Storage)

	}
}

func handleConnection(conn net.Conn, Storage *store.Storage) {
	defer func() {
		conn.Close()
	}()

	for {
		buf := make([]byte, 1024)

		len, err := conn.Read(buf)

		if err != nil {
			conn.Close()
			continue
		}
		SlimBuf := parser.WhiteSpaceTrimmer(string(buf[:len]))
		DecodedBuf := parser.BulkDecoder(SlimBuf)
		Resp := parser.CommandChecker(Storage, DecodedBuf)
		_, err = conn.Write([]byte(Resp))
		if err != nil {
			conn.Close()
			continue
		}

	}
}
