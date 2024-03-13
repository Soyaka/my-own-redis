package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	OK string = "+OK\r\n"
)

func MakeServer() *ServerCred {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		id := uuid.New()
		return NewServerCred("127.0.0.1", "6379", id.String(), "master", "", "")
	case 2:
		if strings.ToLower(args[0]) == "--port" {
			port, err := strconv.Atoi(args[1])
			if err != nil {
				return nil
			}
			id := uuid.New()
			return NewServerCred("127.0.0.1", strconv.Itoa(port), id.String(), "master", "", "")
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
				return NewServerCred("127.0.0.1", strconv.Itoa(port), id.String(), "slave", args[3], strconv.Itoa(ParentPort))
			}
		}
	}
	return nil

}

func StratServer() *ServerCred {
	server := MakeServer()
	if server == nil {
		return nil
	}
	if server.Role == "slave" {
		go server.ConnectReplica()
	}

	return server
}

func (S *ServerCred) ConnectReplica() {
	conn, err := net.Dial("tcp", S.Parent+":"+S.PrntPort)
	if err != nil {
		fmt.Println("#ConnectReplica:faled to connect")
	}
	defer conn.Close()
	_, err = conn.Write([]byte("*1\r\n$4\r\nping\r\n"))
	if err != nil {
		fmt.Println("#ConnectReplica:faled to  ping")
	}
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("#ConnectReplica:faled to read pong")
	}
	fmt.Printf("%s", string(buf))
	_, err = S.SendRelconfPort(conn)
	if err != nil {
		fmt.Println("#ConnectReplica:faled send port")
	}
	_, err = S.SendRelconfCapa(conn)
	if err != nil {
		fmt.Println("#ConnectReplica:failed to send capa")
	}

}

func (S *ServerCred) SendRelconfPort(conn net.Conn) ([]byte, error) {
	buff := make([]byte, 1024)

	msg := fmt.Sprint("*3\r\n$8\r\nREPLCONF\r\n$14\r\nlistening-port\r\n$4\r\n", S.Port, "\r\n")
	_, err := conn.Write([]byte(msg))
	if err != nil {
		return buff, err
	}

	_, err = conn.Read(buff)
	if err != nil {
		return buff, err
	}

	return buff, nil

}

func (S *ServerCred) SendRelconfCapa(conn net.Conn) ([]byte, error) {
	buff := make([]byte, 1024)
	msg := []byte("*3\r\n$8\r\nREPLCONF\r\n$4\r\ncapa\r\n$6\r\npsync2\r\n")
	_, err := conn.Write(msg)
	if err != nil {
		return buff, err
	}

	_, err = conn.Read(buff)
	if err != nil {
		return buff, err
	}

	return buff, nil
}
