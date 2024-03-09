package benchs_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sYamaz/benchmark/benchs"
)

func BenchmarkJsonEncode(b *testing.B) {
	b.Run("marshall/byte配列を取得する", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(benchs.EncodeJsonObject)
			if err != nil {
				b.Error(err)
			}
		}
	})
	b.Run("marshall/writerに書き込む", func(b *testing.B) {
		writer := bytes.NewBuffer([]byte{})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			byteArray, err := json.Marshal(benchs.EncodeJsonObject)
			if err != nil {
				b.Error(err)
			}
			writer.Write(byteArray)
		}
	})
	b.Run("encoder/byte配列を取得する", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			byteArray := []byte{}
			buf := bytes.NewBuffer(byteArray)
			if err := json.NewEncoder(buf).Encode(benchs.EncodeJsonObject); err != nil {
				b.Error(err)
			}
		}
	})
	b.Run("encoder/既存のwriterに書き込む", func(b *testing.B) {
		writer := bytes.NewBuffer([]byte{})

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := json.NewEncoder(writer).Encode(benchs.EncodeJsonObject); err != nil {
				b.Error(err)
			}
		}
	})
}
