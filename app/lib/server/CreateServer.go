package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Server interface {

}

type Kernel struct{
	Server Server
}

func (k *Kernel) Serve() {
	k.Server = MakeServer()
	if k.Server != nil {
		StartServer(k.Server)
	}
}

func (k *Kernel) GetServerInfo() {
	if k.Server == nil {
		fmt.Println("Invalid arguments. Please use --port <port number>")
	}

}












func StartServer(server Server) {
	listener, err := net.Listen(TCP, Port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening on:", Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// Implement logic to handle connections
		 go connect.HandleConnection(conn, Storage, Port)
	}
}








func MakeServer() Server {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		id := uuid.New()
		return NewMasterServer("localhost", "6379", id.String(), "master")
	case 2:
		if strings.ToLower(args[0]) == "--port" {
			port, err := strconv.Atoi(args[1])
			if err != nil {
				return nil
			}
			id := uuid.New()
			return NewMasterServer("localhost", strconv.Itoa(port), id.String(), "master")
		}
	case 5:
		if strings.ToLower(args[0]) == "--port" {
			port, err := strconv.Atoi(args[1])
			if err != nil {
				return nil
			}
			id := uuid.New()
			if strings.ToLower(args[2]) == "--replicaof" {
				ParentPort, err := strconv.Atoi(args[4])
				if err != nil {
					return nil
				}
				return NewReplicaServer("localhost", args[3], strconv.Itoa(ParentPort), strconv.Itoa(port), id.String(), "slave")
			}
		}
	}
	return nil

}
