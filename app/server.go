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

func handleConnection(conn net.Conn) { //1
	defer func() {
		conn.Close()
	}()

	for {
		buf := make([]byte, 1024)

		len, err := conn.Read(buf)

		if err != nil {
			conn.Close()
			continue
		}
		input := buf[:len]

		response := handleCommand(handleDecode(string(input)))
		_, err = conn.Write([]byte(response))
		if err != nil {
			conn.Close()
			break
		}

	}
}

func handleDecode(buff string) []string {
	args := strings.Fields(buff)
	if buff == "ping\n" || buff == "*1\r\n\r\nping\r\n\n" {
		//return "+PONG\r\n"
		retrn := make([]string, 1)
		retrn[0] = "PING"
		return retrn

	}else if strings.Contains(args[0], "*") {
		args = args[1:]
	}
	
	return args
}

func handleCommand(elements []string) string { //3
	switch strings.ToLower(elements[0]) {
	case "ping":
		return "+PONG\r\n"

	case "echo":
		return EncodeResponse(elements[1:])

	}
	return "-err"

}

func EncodeResponse(resSlice []string) (resString string) { //4
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
