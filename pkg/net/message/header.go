package message

import (
	"encoding/binary"
	"hash/crc32"
)

type HeadInfo struct {
	ControlSize uint16
	DataSize    uint32
	HeadCrc32   uint32
}

func (h *HeadInfo) Len() int {
	return 2 + 4 + 4
}

func (h *HeadInfo) SerializeTo(buf []byte) bool {
	binary.LittleEndian.PutUint16(buf, h.ControlSize)
	binary.LittleEndian.PutUint32(buf[2:], h.DataSize)

	crc := h.CRC32(buf[:6])
	binary.LittleEndian.PutUint32(buf, crc)
	return true
}

func (h *HeadInfo) DeserializeFrom(data []byte) bool {
	h.ControlSize = binary.LittleEndian.Uint16(data)
	h.DataSize = binary.LittleEndian.Uint32(data[2:])
	h.HeadCrc32 = binary.LittleEndian.Uint32(data[6:])

	crc := h.CRC32(data[:6])
	return crc == h.HeadCrc32
}

func (h *HeadInfo) CRC32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
