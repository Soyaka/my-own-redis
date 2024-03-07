package conexion

import (
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/lib/cmd"
	"github.com/codecrafters-io/redis-starter-go/app/lib/parser"
	"github.com/codecrafters-io/redis-starter-go/app/lib/server"
	store "github.com/codecrafters-io/redis-starter-go/app/lib/storage"
)

func HandleConnection(conn net.Conn, Storage *store.Storage, server *server.ServerCred) {
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
		SlimBuf := parser.WhiteSpaceTrimmer(string(buf[:len]))
		DecodedBuf := parser.BulkDecoder(SlimBuf)
		Resp := cmd.CommandChecker(Storage, DecodedBuf, server)
		_, err = conn.Write([]byte(Resp))
		if err != nil {
			conn.Close()
			continue
		}

	}
}
