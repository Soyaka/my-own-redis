package server

import (
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Server *ServerCred

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
	return server
}
