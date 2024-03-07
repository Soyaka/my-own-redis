package cmd

import (
	"fmt"
	"strings"

	store "github.com/codecrafters-io/redis-starter-go/app/lib/storage"
)

func CommandChecker(s *store.Storage, elements []string, role string) string {
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
		response = handleInfo(elements, role)
	}

	return response
}
