package cmd

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/lib/server"
	store "github.com/codecrafters-io/redis-starter-go/app/lib/storage"
)

func CommandChecker(s *store.Storage, elements []string, server *server.ServerCred) string {
	var response string
	switch strings.ToUpper(elements[0]) {
	case PING:
		response = fmt.Sprint(PONG)
	case ECHO:
		response = handleECHO(elements)
	case SET:
		response = handleSET(s, elements)
	case GET:
		response = handleGET(s, elements)
	case INFO:
		response = handleInfo(elements, server)
	case REPLCONF:
		response = OK
	case PSYNC:
		response = handlePsync(server)
	}

	return response
}
