package main

import (
	"encoding/binary"
	"fmt"
)

func CheckSlice(buf []byte) int {
	binary.LittleEndian.PutUint16(buf, 123)
	binary.LittleEndian.PutUint32(buf[2:], 123456)

	test := "This is a test. hahahaha!!!"
	l := copy(buf[6:], []byte(test))
	fmt.Println("copy len=", l)
	return 6 + l
}

func main() {
	data := make([]byte, 1024)
	len := CheckSlice(data)
	fmt.Println("data len=", len)
	fmt.Println(string(data))
}
