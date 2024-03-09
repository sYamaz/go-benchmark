package benchs

import (
	"bytes"
	"encoding/json"
	"io"
)

type (
	JsonStructField struct {
		Int    int     `json:"int"`
		Bool   bool    `json:"bool"`
		String string  `json:"string"`
		Float  float64 `json:"float"`
	}

	JsonObject struct {
		Int         int             `json:"int"`
		Bool        bool            `json:"bool"`
		String      string          `json:"string"`
		Float       float64         `json:"float"`
		StringArray []string        `json:"stringArray"`
		Struct      JsonStructField `json:"struct"`
	}
)

var (
	EncodeJsonObject = JsonObject{
		Int:    100,
		Bool:   true,
		String: "this is string field",
		Float:  200.0058,
		StringArray: []string{
			"String A",
			"String B",
			"String C",
			"String D",
			"String E",
			"String F",
			"String G",
			"String H",
			"String I",
			"String J",
			"String K",
			"String L",
			"String M",
			"String N",
		},
		Struct: JsonStructField{
			Int:    -10,
			Bool:   false,
			String: "this is string field",
			Float:  -1200.998,
		},
	}

	EncodedJsonAsBytes  = getEncodedJsonAsBytes()
	EncodedJsonAsString = string(EncodedJsonAsBytes)
)

func getEncodedJsonAsBytes() []byte {
	b, _ := json.Marshal(EncodeJsonObject)
	return b
}

func EncodedJsonAsNewReader() io.Reader {
	return bytes.NewReader(EncodedJsonAsBytes)
}
