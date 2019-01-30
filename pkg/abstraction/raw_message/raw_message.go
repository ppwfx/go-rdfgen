package main

import "encoding/json"

type metaTrait struct {
	Id string `json:"@id,omitempty"`
	Type string `json:"@type,omitempty"`
}

type Review struct {
	Item interface{}
}

func Decode(b []byte) interface{} {
	meta := metaTrait{}

	json.Unmarshal(b, meta)
}
