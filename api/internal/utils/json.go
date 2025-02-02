package utils

import (
	"encoding/json"
	"io"
)

func ParseJson[T any](reader io.Reader) (*T, error) {
	var data T
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
