// 使用Binary包序列化的限制：
// "Data must be a fixed-size value or a slice of fixed-size value."
// 如果字段中有不确定大小的类型（如 int, string等），编译没问题，但会在运行时反序列化报错:
// "binary write error:binary.Write: invalid type main.Message"
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

// OK
type Message1 struct {
	Tag    uint64
	Length uint64
	Value  [3]uint8
}

// Error
/*
type Message1 struct {
	Id    int
	Size  int
	Value []uint8
}
*/

func BinaryRw() {
	m1 := Message1{Tag: 1, Length: 3, Value: [3]uint8{3, 4, 5}}
	//m1 := Message1{Tag: 1, Length: 3, Value: make([]uint8, 3)} // Run Error
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, m1); err != nil {
		log.Fatal("binary write error: ", err)
	}
	fmt.Printf("m1: tag=%d, length=%d, value=%v\n", m1.Tag, m1.Length, m1.Value)

	bufLen := buf.Len()
	fmt.Printf("buf.Len()=%d\n", bufLen)
	var i uint8
	for ch, err := buf.ReadByte(); err == nil; ch, err = buf.ReadByte() {
		fmt.Printf("[%d] 0x%x\n", i, ch)
		i++
	}

	/*
		var m2 Message1
		if err := binary.Read(buf, binary.LittleEndian, &m2); err != nil {
			log.Fatal("binary read error: ", err)
		}
		fmt.Printf("m2: tag=%d, length=%d, value=%v\n", m2.Tag, m2.Length, m2.Value)
	*/
}

func main() {
	BinaryRw()
}
