package main

import (
	"fmt"
	"net"

	connect "github.com/codecrafters-io/redis-starter-go/app/lib/conexion"
	"github.com/codecrafters-io/redis-starter-go/app/lib/server"
	store "github.com/codecrafters-io/redis-starter-go/app/lib/storage"
)

const (
	TCP  = "tcp"
	PORT = ":6379"
)

func main() {
	server := server.StratServer()
	Port := ":" + server.Port
	Role := server.Role
	listener, err := net.Listen(TCP, Port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	Storage := store.NewStorage()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go connect.HandleConnection(conn, Storage, Role)
	}
}
