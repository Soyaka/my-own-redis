package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/lib/server"
	store "github.com/codecrafters-io/redis-starter-go/app/lib/storage"
)

const (
	SEPARATOR string = "\r\n"
	PLUS      string = "+"
	MINUS     string = "-"
	DOLLAR    string = "$"
	STAR      string = "*"
	PING      string = "PING"
	PONG      string = "+PONG\r\n"
	ECHO      string = "ECHO"
	SET       string = "SET"
	INFO      string = "INFO"
	GET       string = "GET"
	OK        string = "+OK\r\n"
	NON       string = "$-1\r\n"
)

func handleECHO(elements []string) string {
	var response string
	if len(elements) > 2 {
		return fmt.Sprint(STAR, len(elements)-1, SEPARATOR)
	}
	for _, element := range elements[1:] {
		response += fmt.Sprint(DOLLAR, len(element), SEPARATOR, element, SEPARATOR)
	}
	return response

}

func handleGET(s *store.Storage, args []string) string {
	key := args[1]
	value, ok := s.GetValue(key)
	if !ok {
		return NON
	}
	return fmt.Sprint(DOLLAR, len(value), SEPARATOR, value, SEPARATOR)
}

func handleSET(s *store.Storage, args []string) string {
	len := len(args)
	switch len {
	case 3:
		if err := handleSETWXP(s, args); err != nil {
			return NON
		}
		return OK
	case 5:
		if err := handleSETXP(s, args); err != nil {
			return NON
		}
		return OK
	}
	return NON

}

func handleSETWXP(s *store.Storage, args []string) error {
	var expTime time.Time
	var data *store.Data
	key := args[1]
	expTime = time.Now().Add(999999 * time.Hour)
	data = store.NewData(args[2], expTime)
	s.SetValue(key, data)
	return nil
}

func handleSETXP(s *store.Storage, args []string) error {
	var expTime time.Time
	var data *store.Data
	key := args[1]
	timeXP, err := strconv.Atoi(args[4])
	if err != nil {
		return err
	}
	expTime = time.Now().Add(time.Duration(timeXP) * time.Millisecond)
	data = store.NewData(args[2], expTime)
	s.SetValue(key, data)
	return nil
}

func handleInfo(args []string, server *server.ServerCred) string {
	var rsSlice []string
	var response string
	switch strings.ToUpper(args[1]) {
	case "REPLICATION":
		if server.Role == "master" {
			rsSlice = append(rsSlice, "role:master")
			rsSlice = append(rsSlice, "master_replid:8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb")
			rsSlice = append(rsSlice, "master_repl_offset:0")
		} else {
			rsSlice = append(rsSlice, "role:slave")
		}
	}
	for _, resp := range rsSlice {
		response += fmt.Sprint(DOLLAR, len(resp), SEPARATOR, resp, SEPARATOR)
	}
	// if len(rsSlice) > 1 {
	// 	response = fmt.Sprint(STAR,len(rsSlice),SEPARATOR,response)
	// }

	return response
}
