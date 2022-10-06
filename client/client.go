package client

type SERVER_TYPE int8

const (
	SERVER_HTTP SERVER_TYPE = 0
	SERVER_TCP  SERVER_TYPE = 1
)

type COMMAND_TYPE int8

const (
	COMMAND_GET COMMAND_TYPE = 0
	COMMAND_SET COMMAND_TYPE = 1
	COMMAND_DEL COMMAND_TYPE = 2
)

type Cmd struct {
	Name  COMMAND_TYPE
	Key   string
	Value string
	Error error
}

type Client interface {
	Run(*Cmd)
	PipelinedRun([]*Cmd)
}

func New(typ SERVER_TYPE, addr string) Client {
	if typ == SERVER_HTTP {
		return newHTTPClient(addr)
	}
	if typ == SERVER_TCP {
		return newTCPClient(addr)
	}
	panic("unknown client type")
}
