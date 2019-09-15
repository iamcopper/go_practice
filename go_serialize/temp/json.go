package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message3 struct {
	Tag    uint64  `json:"tag"`
	Length uint64  `json:"size"`
	Value  []uint8 `json:"value"`
}

func JsonEncodeDecode() {
	m1 := Message3{Tag: 3, Length: 3, Value: make([]uint8, 3)}
	m1.Value[0], m1.Value[1], m1.Value[2] = 7, 8, 9

	var buf []byte
	var err error

	if buf, err = json.Marshal(m1); err != nil {
		log.Fatal("json marshal error:", err)
	}
	fmt.Printf("m1: tag=%d, length=%d, value=%v\n", m1.Tag, m1.Length, m1.Value)

	fmt.Printf("bufLen=%d\n", len(buf))
	for i := 0; i < len(buf); i++ {
		fmt.Printf("[%d] 0x%x\n", i, buf[i])
	}

	var m2 Message3
	if err = json.Unmarshal(buf, &m2); err != nil {
		log.Fatal("json unmarahsl error:", err)
	}
	fmt.Printf("m2: tag=%d, length=%d, value=%v\n", m2.Tag, m2.Length, m2.Value)
}

func main() {
	JsonEncodeDecode()
}
