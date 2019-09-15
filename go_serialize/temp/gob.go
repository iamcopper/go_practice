package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Message2 struct {
	Tag    uint64
	Length uint64
	Value  []uint8
}

func GobEncodeDecode() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	m1 := Message2{Tag: 2, Length: 3, Value: []uint8{1, 2, 3}}
	if err := enc.Encode(m1); err != nil {
		log.Fatal("encode error:", err)
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
		dec := gob.NewDecoder(&buf)

		var m2 Message2
		if err := dec.Decode(&m2); err != nil {
			log.Fatal("decode error:", err)
		}
		fmt.Printf("m2: tag=%d, length=%d, value=%v\n", m2.Tag, m2.Length, m2.Value)
	*/
}

func main() {
	GobEncodeDecode()
}
