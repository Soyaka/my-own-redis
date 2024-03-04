package parser

import (
	"fmt"
	"strings"
)

const (
	SEPARATOR = "\r\n"
	PLUS      = "+"
	MINUS     = "-"
	DOLLAR    = "$"
	STAR      = "*"
	PING      = "PING"
	PONG      = "PONG"
	ECHO      = "ECHO"
	ERR       = "INVALID"
)

//1

func WhiteSpaceTrimmer(str string) []string {
	args := strings.Fields(str)
	return args
}

// 2
func BulkDecoder(args []string) []string {
	var elements []string
	for _, arg := range args {
		if strings.Contains(arg, STAR) || strings.Contains(arg, DOLLAR) {
			continue
		} else {
			elements = append(elements, arg)
		}
	}
	return elements
}

//3

func CommandChecker(elements []string) string {
	switch strings.ToUpper(elements[0]) {
	case PING:
		return fmt.Sprint(PLUS, PONG, SEPARATOR)
	case ECHO:
		var response string 
		if len(elements) > 2 {
			response = fmt.Sprint(STAR, len(elements)-1, SEPARATOR)
		}
		for _, element := range elements[1:] {
			response += fmt.Sprint(DOLLAR, len(element), SEPARATOR, element, SEPARATOR)
			return response
		}
	}
	return ERR
}