package common

import "errors"

const (
	InvalidLocalIP = "invalid local ip"
	InvalidTcpConn = "invalid tcp conn"
)

var (
	ErrSendBufferIsFull = errors.New("send buffer full")
)
