package helper

import "strings"

var typeMap = map[string]string{}

func init() {
	typeMap["Text"] = "string"
	typeMap["URL"] = "string"
	typeMap["Number"] = "float64"
	typeMap["Boolean"] = "bool"
}

type SchemaOrg struct {
}

func (h SchemaOrg) CanToUrl(id string) (bool) {
	is := strings.Contains(id, "schema.org")
	return is
}

func (h SchemaOrg) ToUrl(id string) (url string) {
	if strings.HasSuffix(id, ".jsonld") {
		return id
	}

	url = id + ".jsonld"

	return url
}

func (h SchemaOrg) CanMapType(t string) (bool) {
	is := strings.Contains(t, "schema.org")
	return is
}

func (h SchemaOrg) MapType(t string) (string) {
	t = strings.Replace(t, "http://schema.org/", "", -1)

	_, ok := typeMap[t]
	if ok {
		t = typeMap[t]
	}

	return t
}