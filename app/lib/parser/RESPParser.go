package parser

import (
	"strings"
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