package types

import (
	"net/http"
	"encoding/gob"
	"strings"
)

func init() {
	gob.Register(Response{})
}

type ContentType string

const (
	TEXT_TURTLE ContentType = "text/turtle"
	JSON_LD     ContentType = "application/ld+json"
)

type Response struct {
	Body []byte
	Header http.Header
}

type Triple struct {
	Subject string
	Predicate string
	Object string
}

type IdExpander interface {
	CanToUrl(string) (bool)
	ToUrl(string) (string)
}

type Parser interface {
	CanParse(response Response) (bool)
	Parse([]byte) (Triples, error)
}

type TypeMapper interface {
	CanMapType(string) (bool)
	MapType(string) (string)
}

type Triples []Triple

func (_triples Triples) RemoveDuplicates() (triples Triples) {
	m := map[string]bool{}

	for _, t := range  _triples {
		s := t.Subject + t.Predicate + t.Object
		_, ok := m[s]
		if ok  {
			continue
		}

		triples = append(triples, t)

		m[s]= true
	}

	return
}

func (_triples Triples) FilterBySubject(subject string) (triples Triples) {
	for _, t := range  _triples {
		if t.Subject != subject {
			continue
		}

		triples = append(triples, t)
	}

	return
}

func (_triples Triples) FilterBySubjectContains(contains string) (triples Triples) {
	for _, t := range  _triples {
		if !strings.Contains(t.Subject, contains) {
			continue
		}

		triples = append(triples, t)
	}

	return
}

func (_triples Triples) FilterByPredicate(predicate string) (triples Triples) {
	for _, t := range  _triples {
		if t.Predicate != predicate {
			continue
		}

		triples = append(triples, t)
	}

	return
}

func (_triples Triples) FilterByObject(object string) (triples Triples) {
	for _, t := range  _triples {
		if t.Object != object {
			continue
		}

		triples = append(triples, t)
	}

	return
}

func (triples Triples) MapBySubject() (g map[string]map[string][]string) {
	g = map[string]map[string][]string{}

	for _, t := range triples {
		_, k := g[t.Subject]
		if !k {
			g[t.Subject] = map[string][]string{}
		}

		_, k = g[t.Subject][t.Predicate]
		if !k {
			g[t.Subject][t.Predicate] = []string{}
		}

		g[t.Subject][t.Predicate] = append(g[t.Subject][t.Predicate], t.Object)
	}

	return
}

func (triples Triples) GroupBySubject() (g map[string][]Triple) {
	g = map[string][]Triple{}

	for _, t := range triples {
		_, k := g[t.Subject]
		if !k {
			g[t.Subject] = []Triple{}
		}

		g[t.Subject] = append(g[t.Subject], t)
	}

	return
}

func (triples Triples) GroupByPredicate() (g map[string][]Triple) {
	g = map[string][]Triple{}

	for _, t := range triples {
		_, k := g[t.Predicate]
		if !k {
			g[t.Predicate] = []Triple{}
		}

		g[t.Predicate] = append(g[t.Predicate], t)
	}

	return
}

func (triples Triples) GroupByObject() (g map[string][]Triple) {
	g = map[string][]Triple{}

	for _, t := range triples {
		_, k := g[t.Object]
		if !k {
			g[t.Object] = []Triple{}
		}

		g[t.Object] = append(g[t.Object], t)
	}

	return
}