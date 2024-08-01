package network

type ClientIf interface {
	Init(string, string, string) error
	Server() error
	Stop() error

	Write([]byte) (int, error)
	Read() ([]byte, error)
}

type ServerIf interface {
	Init() error
	Server() error
	Stop() error

	Write([]byte) (int, error)
	Read() ([]byte, error)
}
