package record

import (
	"bytes"
	"fmt"

	"github.com/linkedin/goavro/v2"
)

func FormatAvroToMap(data []byte) ([]map[string]interface{}, error) {
	r, err := goavro.NewOCFReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	maps := make([]map[string]interface{}, 0)
	for r.Scan() {
		v, err := r.Read()
		if err != nil {
			return nil, err
		}
		m, mapOk := v.(map[string]interface{})
		if !mapOk {
			return nil, fmt.Errorf("invalid input: %v", v)
		}
		maps = append(maps, m)
	}

	return maps, nil
}