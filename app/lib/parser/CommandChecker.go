package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/lib/store"
)

func CommandChecker(s *store.Storage, elements []string) string {
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
	}
	return response
}

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