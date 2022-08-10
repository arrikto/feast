package model

import (
	"encoding/json"
	"errors"
)

func MapStrEncode(map_ map[string]string) ([]byte, error) {
	b, err := json.Marshal(map_)
	if err != nil {
		return nil, errors.New("couldn't convert map[string][string] to []byte")
	}
	return b, nil
}

func MapStrDecode(b []byte) (map[string]string, error) {
	var map_ map[string]string
	err := json.Unmarshal(b, &map_)
	if err != nil {
		return nil, errors.New("couldn't convert []byte to map[string][string]")
	}
	return map_, nil
}
