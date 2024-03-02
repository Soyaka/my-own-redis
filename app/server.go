package main

import (
	"fmt"
	"net"
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
	defer func() {
		conn.Close()
	}()
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			continue
		}
		response := handleCommand(handleDecode(string(buf)))
		_, err = conn.Write([]byte(response))
		if err != nil {
			conn.Close()
			break
		}

	}
}

func handleDecode(buff string) []string {
	args := strings.Split(buff, "\r\n")
	netArgs := []string{}
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "*") || strings.HasPrefix(args[i], "$") {
			i++
		} else {
			netArgs = append(netArgs, args[i])
		}
	}
	return netArgs
}

func handleCommand(elements []string) string {
	switch strings.ToLower(elements[0]) {
	case "echo":
		return strings.Join(elements[1:], " ")
	case "ping":
		return "+PONG\r\n"
	}
	return ""
}
