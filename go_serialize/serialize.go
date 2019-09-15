package serialize

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"log"
)

type Message1 struct {
	Tag    uint64
	Length uint64
	Value  [3]uint8
}

type Message2 struct {
	Tag    uint64  `json:"tag"`
	Length uint64  `json:"length"`
	Value  []uint8 `json:"value"`
}

func BinaryRW() {
	m1 := Message1{Tag: 1, Length: 3, Value: [3]uint8{3, 4, 5}}
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, m1); err != nil {
		log.Fatal("binary write error: ", err)
	}

	var m2 Message1
	if err := binary.Read(buf, binary.LittleEndian, &m2); err != nil {
		log.Fatal("binary read error: ", err)
	}
}

func GobEncodeDecode() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	m1 := Message1{Tag: 2, Length: 3, Value: [...]uint8{6, 7, 8}}
	if err := enc.Encode(m1); err != nil {
		log.Fatal("encode error:", err)
	}

	var m2 Message1
	if err := dec.Decode(&m2); err != nil {
		log.Fatal("decode error:", err)
	}
}

func JsonEncodeDecode() {
	m1 := Message2{Tag: 3, Length: 3, Value: make([]uint8, 3)}

	var buf []byte
	var err error

	if buf, err = json.Marshal(m1); err != nil {
		log.Fatal("json marshal error:", err)
	}

	var m2 Message2
	if err = json.Unmarshal(buf, &m2); err != nil {
		log.Fatal("json unmarahsl error:", err)
	}
}
