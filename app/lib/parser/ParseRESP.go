package parser

import (
	"strings"
)

const (
	DOLLAR string = "$"
	STAR   string = "*"
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
