package serialize

import (
	"testing"
)

func BenchmarkBinaryRW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinaryRW()
	}
}

func BenchmarkGobEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GobEncodeDecode()
	}
}

func BenchmarkJsonEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonEncodeDecode()
	}
}
