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
	PING        string = "*1\r\n$4\r\nping\r\n"
	PONG        string = "+PONG\r\n"
	OK          string = "+OK\r\n"
	RPLCONFPORT string = "*3\r\n$8\r\nREPLCONF\r\n$14\r\nlistening-port\r\n$4\r\n6380\r\n"
	RPLCONFCAPA string = "*3\r\n$8\r\nREPLCONF\r\n$4\r\ncapa\r\n$6\r\npsync2\r\n"
	PSYNC       string = "*3\r\n$5\r\nPSYNC\r\n$1\r\n?\r\n$2\r\n-1\r\n"
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

//TODO: PUT ALL THE CODE BELLOW IN A SPECIFIC PACKAGE FOR REPLICA

func (S ServerCred) ConnectReplica() {
	msgList := []string{
		PING,
		RPLCONFPORT,
		RPLCONFCAPA,
		PSYNC,
	}
	msgMap := map[string]string{
		PING:        PONG,
		RPLCONFPORT: OK,
		RPLCONFCAPA: OK,
		PSYNC:       OK,
	}

	conn, err := net.Dial("tcp", S.Parent+":"+S.PrntPort)
	if err != nil {
		fmt.Println("#ConnectReplica:failed to connect")
	}
	defer conn.Close()
	for _, msg := range msgList {
		resp, err := S.SendHandshake(conn, msg)
		if err != nil {
			fmt.Println("#ConnectReplica: err ", err)
			break
		}
		if string(resp) != msgMap[msg] {
			fmt.Println("#ConnectReplica: response dows not match", string(resp))
		} else {
			fmt.Println(string(resp))
		}
	}

}

func (S ServerCred) SendHandshake(conn net.Conn, msg string) ([]byte, error) {
	buff := make([]byte, 1024)
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
