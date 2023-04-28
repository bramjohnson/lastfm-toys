package util

import "encoding/xml"

func SafeUnmarshal[T any](data []byte, _ T) *T {
	var unmarshalled T
	if err := xml.Unmarshal(data, &unmarshalled); err != nil {
		panic(err)
	}
	return &unmarshalled
}
