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
		return fmt.Sprint(PONG)
	case ECHO:

		if len(elements) > 2 {
			response = fmt.Sprint(STAR, len(elements)-1, SEPARATOR)
		}
		for _, element := range elements[1:] {
			response += fmt.Sprint(DOLLAR, len(element), SEPARATOR, element, SEPARATOR)
		}
	case SET:
		SETlen := len(elements)
		switch SETlen {
		case 3:

		}
		//TODO: //add the SET logic
		// _expire, err := strconv.Atoi(elements[4])
		// if err != nil {
		// 	response = NON
		// }
		// expireDuration := time.Duration(_expire) * time.Microsecond
		// data := store.Data{
		// 	Value:    elements[2],
		// 	ExpireAt: time.Now().Add(expireDuration),
		// }
		// s.SetValue(elements[1], data)
		// response = OK
		response = handleSET(s, elements)
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

//FIXME: fix the bew function : set appropriate time

func handleSET(s *store.Storage, args []string) string {
	var expirationTime time.Time

	var key string = args[1]
	var value string = args[2]
	var data = store.NewData(value, expirationTime)

	len := len(args)
	switch {
	case len == 3:
		expirationTime = time.Now().Add(999999 * time.Hour)
		data.ExpriredAt = expirationTime
		s.SetValue(key, data)
	case len == 5:
		_expire, err := strconv.Atoi(args[4])
		if err != nil {
			return NON
		}
		expireDuration := time.Duration(_expire) * time.Microsecond
		expirationTime := time.Now().Add(expireDuration)
		data.ExpriredAt = expirationTime

		s.SetValue(key, data)
	case len > 5 || len < 3:
		return NON
	}
	return OK

}
