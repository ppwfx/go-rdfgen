package helper

import "strings"

type PurlOrg struct {
}

func (h PurlOrg) CanToUrl(id string) (bool) {
	is := strings.Contains(id, "purl.org")
	return is
}

func (h PurlOrg) ToUrl(id string) (url string) {
	if strings.Contains(id, "http://purl.org/dc/terms") {
		return "http://purl.org/dc/terms"
	}

	if strings.Contains(id, "http://purl.org/dc/elements/1.1") {
		return "http://purl.org/dc/elements/1.1"
	}

	if strings.Contains(id, "http://purl.org/dc/dcmitype") {
		return "http://purl.org/dc/dcmitype"
	}

	if strings.Contains(id, "http://purl.org/dc/dcam/") {
		return "http://purl.org/dc/dcam/"
	}

	return id
}
