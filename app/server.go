package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :6379...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Accepted connection from:", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()
		if strings.ToUpper(command) == "PING" {
			response := "+PONG\r\n"
			conn.Write([]byte(response))
			fmt.Printf("Received: %s, Sent: %s", command, response)
		}
	}
}












type RESPParser struct{}

func (RS *RESPParser) stringParser(RespString string) (string, string) {
	if strings.HasPrefix(RespString, "$") {
		parts := strings.SplitN(RespString, "\r\n", 2)
		_, err := strconv.Atoi(parts[0])
		if err!= nil{
			return "", RespString
		}

		data:= parts[1]
		fmt.Println(data)

	}
	

	return "", ""

}
