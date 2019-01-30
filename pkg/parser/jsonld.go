package parser

import (
	"strings"
	"github.com/21stio/go-rdfgen/pkg/types"
	"encoding/json"
	"github.com/piprate/json-gold/ld"
	"errors"
	"regexp"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

var splitRegex = regexp.MustCompile(`(<([a-zA-Z0-9/.:#-_]*)>|_:[a-z0-9]*) <([a-zA-Z0-9/.:#-_]*)> ("(.*)"|<([a-zA-Z0-9/.:#-_]*)>|_:[a-z0-9]*)`)
var proc = ld.NewJsonLdProcessor()
var options = ld.NewJsonLdOptions("")

func init() {
	options.Format = "application/nquads"
	options.Algorithm = "URDNA2015"
	options.DocumentLoader = ld.NewCachingDocumentLoader(ld.NewDefaultDocumentLoader(&http.Client{}))
}

type JsonLd struct {
}

func (p JsonLd) CanParse(r types.Response) (bool) {
	if strings.Contains(r.Header.Get(CONTENT_TYPE), string(types.JSON_LD)) {
		return true
	}

	return false
}

func (p JsonLd) Parse(b []byte) (triples types.Triples, err error) {
	d := map[string]interface{}{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return
	}

	normalizedTriples, err := proc.Normalize(d, options)
	if err != nil {
		return
	}

	s, ok := normalizedTriples.(string)
	if !ok {
		err = errors.New("not a string")
	}

	for _, s := range strings.Split(s, "\n") {
		if s == "" {
			break
		}

		m := splitRegex.FindAllStringSubmatch(s, -1)

		if len(m) == 0 {
			spew.Dump(s)
			continue
		}

		triple := types.Triple{}
		triple.Subject = getSubject(m[0])
		triple.Predicate = getPredicate(m[0])
		triple.Object = getObject(m[0])

		triples = append(triples, triple)
	}

	return
}

func getSubject(m []string) (string) {
	if m[2] != "" {
		return m[2]
	}

	return m[1]
}

func getPredicate(m []string) (string) {
	return m[3]
}

func getObject(m []string) (string) {
	if m[6] != "" {
		return m[6]
	}
	if m[5] != "" {
		return m[5]
	}

	return m[4]
}