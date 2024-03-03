package parser

import (
	"fmt"
	"strings"
)

func HandleDecode(buff string) []string {
	args := strings.Fields(buff)
	if strings.Contains(args[0], "*") {
		args = args[1:]
	}
	return args
}

func HandleCommand(elements []string) string {
	switch strings.ToLower(elements[0]) {
	case "echo":

		return EncodeResponse(elements[1:])

	case "ping":
		return "+PONG\r\n"
	}
	return "-ERR"
}

func EncodeResponse(resSlice []string) (resString string) {
	if len(resSlice) == 1 {
		resString = "$" + fmt.Sprint(len(resSlice[0])) + "\r\n" + resSlice[0] + "\r\n"
	} else if len(resSlice) > 1 {
		resString = "*" + fmt.Sprint(len(resSlice)) + "\r\n"
		for _, element := range resSlice {
			resString += "$" + fmt.Sprint(len(element)) + "\r\n" + element + "\r\n"
		}
	}
	return resString
}
