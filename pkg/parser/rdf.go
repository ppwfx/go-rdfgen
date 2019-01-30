package parser

import (
	"strings"
	"github.com/21stio/go-rdfgen/pkg/types"
	"github.com/knakk/rdf"
	"io"
	"bytes"
)

type Rdf struct {
}

func (p Rdf) CanParse(r types.Response) (bool) {
	if strings.Contains(r.Header.Get(CONTENT_TYPE), string(types.TEXT_TURTLE)) {
		return true
	}

	return false
}

func (p Rdf) Parse(b []byte) (triples types.Triples, err error) {
	buf := bytes.NewBuffer(b)

	dec := rdf.NewTripleDecoder(buf, rdf.Turtle)
	for t, err := dec.Decode(); err != io.EOF; t, err = dec.Decode() {
		triple := types.Triple{
			Subject:   t.Subj.String(),
			Predicate: t.Pred.String(),
			Object:    t.Obj.String(),
		}
		triples = append(triples, triple)
	}

	return
}
