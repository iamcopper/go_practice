package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

type Message1 struct {
	Tag    uint64
	Length uint64
	Value  []uint8
}

func BinaryRw() {
	m1 := Message1{Tag: 1, Length: 3, Value: []uint8{13, 14, 15, 15}}
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, m1.Tag); err != nil {
		log.Fatal("binary write error: ", err)
	}
	if err := binary.Write(buf, binary.LittleEndian, m1.Length); err != nil {
		log.Fatal("binary write error: ", err)
	}
	// 序列化切片
	if err := binary.Write(buf, binary.LittleEndian, m1.Value); err != nil {
		log.Fatal("binary write error: ", err)
	}
	fmt.Printf("m1: tag=%d, length=%d, value=%v\n", m1.Tag, m1.Length, m1.Value)

	var m2 Message1
	if err := binary.Read(buf, binary.LittleEndian, &m2.Tag); err != nil {
		log.Fatal("binary read error: ", err)
	}
	if err := binary.Read(buf, binary.LittleEndian, &m2.Length); err != nil {
		log.Fatal("binary read error: ", err)
	}
	// 反序列化切片: 需要先知道切片的长度，然后才能反序列化成功.
	m2.Value = make([]uint8, 4)
	if err := binary.Read(buf, binary.LittleEndian, &m2.Value); err != nil {
		log.Fatal("binary read error: ", err)
	}
	fmt.Printf("m2: tag=%d, length=%d, value=%v\n", m2.Tag, m2.Length, m2.Value)
}

func main() {
	BinaryRw()
}
