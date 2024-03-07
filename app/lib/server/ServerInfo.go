package server

type ServerInfo struct {
	Role string
	Host string
	Port string
	ID   string
}

type ReplicaServer struct {
	*ServerInfo
	Parent     string
	ParentPort string
}

type MasterServer struct {
	*ServerInfo
	Replicas []ReplicaServer
}

func NewReplicaServer(host, parent, parentPort, port, id, role string) *ReplicaServer {
	serverInfo := &ServerInfo{
		Role: role,
		Host: host,
		Port: port,
		ID:   id,
	}
	server := &ReplicaServer{
		ServerInfo: serverInfo,
		Parent:     parent,
		ParentPort: parentPort,
	}
	return server
}

func NewMasterServer(host, port, id, role string) *MasterServer {
	serverInfo := &ServerInfo{
		Role: role,
		Host: host,
		Port: port,
		ID:   id,
	}
	server := &MasterServer{
		ServerInfo: serverInfo,
		Replicas:   make([]ReplicaServer, 0),
	}
	return server
}
