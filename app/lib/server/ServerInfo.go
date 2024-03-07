package server

type ServerCred struct {
	Role string
	Host string
	Port string
	ID   string
	Parent string
	PrntPort string
}

func NewServerCred(host, port, id, role, parent, prntPort string) *ServerCred {
	Server := &ServerCred{
		Role: role,
		Host: host,
		Port: port,
		ID:   id,
		Parent: parent,
		PrntPort: prntPort,
	}
	return Server
}