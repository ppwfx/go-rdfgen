package helper

import (
	"strings"
	"errors"
	"fmt"
)

type HashTag struct {
}

func (h HashTag) CanToUrl(id string) (bool) {
	n := strings.Count(id, "#")
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n > 1 {
		panic(errors.New(fmt.Sprintf("more than one hashtag in id: %v", id)))
	}

	return false
}

func (h HashTag) ToUrl(id string) (url string) {
	return strings.Split(id, "#")[0]
}