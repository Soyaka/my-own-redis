 package parser


 import (
	 "fmt"
	 "net"
	 "strings"
 )

 
 func HandleConnection(conn net.Conn) {
	 defer func() {
		 conn.Close()
	 }()
	 buf := make([]byte, 1024)
 
	 for {
		 len, err := conn.Read(buf)
 
		 if err != nil {
			 conn.Close()
			 continue
		 }
 
		 response := handleCommand(handleDecode(string(buf[0 : len-1])))
		 fmt.Println(response)
		 _, err = conn.Write([]byte(response))
		 if err != nil {
			 conn.Close()
			 break
		 }
 
	 }
 }
 
 
 
 
 func HandleDecode(buff string) []string {
	 args := strings.Fields(buff)
	if len(args) > 2{
		args = args[1:]
	}
	 return args
 }

 
 
 func HandleCommand(elements []string) string {
	 fmt.Println(elements[1])
	 switch strings.ToLower(elements[1]) {
	 case "echo":
		 fmt.Println("+" + strings.Join(elements[1:], "") + "\r\n")
		 return EncodeResponse(elements[2:])
	 case "ping":
		 return "+PONG\r\n"
	 }
	 return "non"
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