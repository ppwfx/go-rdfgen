package helper

import "strings"

type DublinCoreOrg struct {
}

func (h DublinCoreOrg) CanToUrl(id string) (bool) {
	is := strings.Contains(id, "dublincore.org")
	return is
}

func (h DublinCoreOrg) ToUrl(id string) (url string) {
	if strings.Contains(id, "http://dublincore.org/2012/06/14/dcterms") {
		return "http://purl.org/dc/terms"
	}

	return id
}
