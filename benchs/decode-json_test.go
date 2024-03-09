package benchs_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/sYamaz/benchmark/benchs"
)

func BenchmarkJsonDecode(b *testing.B) {
	b.Run("unmarshall/byte配列からデコード", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			got := new(benchs.JsonObject)
			if err := json.Unmarshal(benchs.EncodedJsonAsBytes, got); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("unmarshall/readerからデコード", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			reader := benchs.EncodedJsonAsNewReader()
			b.StartTimer()
			got := new(benchs.JsonObject)
			byteArray, _ := io.ReadAll(reader)
			if err := json.Unmarshal(byteArray, got); err != nil {
				b.Error(err)
			}
		}
	})
	b.Run("decoder/byte配列からデコード", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			got := new(benchs.JsonObject)
			if err := json.NewDecoder(bytes.NewReader(benchs.EncodedJsonAsBytes)).Decode(got); err != nil {
				b.Error(err)
			}
		}
	})
	b.Run("decoder/readerからデコード", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			reader := benchs.EncodedJsonAsNewReader()
			b.StartTimer()
			got := new(benchs.JsonObject)
			if err := json.NewDecoder(reader).Decode(got); err != nil {
				b.Error(err)
			}
		}
	})
}
