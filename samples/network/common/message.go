package common

type MessageHead struct {
	Type      uint16
	Serial    uint16
	Dst       uint16
	Src       uint16
	SessionID uint32
	DataLen   uint32
	DataCrc   uint32
}
