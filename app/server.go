package main

import (
	"fmt"
	"net"

	connect "github.com/codecrafters-io/redis-starter-go/app/lib/conexion"
	store "github.com/codecrafters-io/redis-starter-go/app/lib/storage"
)

const (
	TCP  = "tcp"
	PORT = ":6379"
)

func main() {

	listener, err := net.Listen(TCP, Port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	Storage := store.NewStorage()
	fmt.Println("Listening on :", Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go connect.HandleConnection(conn, Storage, Port)
	}
}
