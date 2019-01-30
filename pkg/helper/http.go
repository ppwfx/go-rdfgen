package helper

import (
	"strings"
)

type Http struct {
}

func (h Http) CanToUrl(id string) (bool) {
	return strings.HasPrefix(id, "https")
}

func (h Http) ToUrl(id string) (url string) {
	return strings.Replace(id, "https", "http", 1)
}
