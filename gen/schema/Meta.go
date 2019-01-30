package schema

type MetaTrait struct {
	Id string `json:"@id,omitempty" jsonld:"@id"`
	Type string `json:"@type,omitempty" jsonld:"@type"`
}