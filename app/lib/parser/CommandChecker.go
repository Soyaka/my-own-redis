package parser

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/lib/store"
)

func CommandChecker(s *store.Storage, elements []string) string {
	var response string
	switch strings.ToUpper(elements[0]) {
	case PING:
		return fmt.Sprint(PONG)
	case ECHO:

		if len(elements) > 2 {
			response = fmt.Sprint(STAR, len(elements)-1, SEPARATOR)
		}
		for _, element := range elements[1:] {
			response += fmt.Sprint(DOLLAR, len(element), SEPARATOR, element, SEPARATOR)

		}
	case SET:
		s.SetValue(elements[1], elements[2])
		response = OK
	case GET:
		value, ok := s.GetValue(elements[1])
		if !ok {
			response = NON
		} else {
			response = fmt.Sprint(DOLLAR, len(value), SEPARATOR, value, SEPARATOR)
		}
	}
	return response
}
