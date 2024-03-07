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
	var response strings.Builder
	for _, resp := range rsSlice {
		response.WriteString(fmt.Sprint(DOLLAR, len(resp), SEPARATOR, resp, SEPARATOR))
	}

	if len(rsSlice) >= 2 {
		response.WriteString(fmt.Sprint(STAR, len(rsSlice), SEPARATOR, response.String(), SEPARATOR))
	}

	return response.String()
}
// remote: [replication-4] Running tests for Replication > Stage #4: Initial Replication ID and Offset
// remote: [replication-4] $ ./spawn_redis_server.sh
// remote: [replication-4] $ redis-cli INFO replication
// remote: [replication-4] Expected: 'master_replid' key in INFO replication.
// remote: [replication-4] Test failed
// remote: [replication-4] Terminating program
// remote: [replication-4] Program terminated successfully