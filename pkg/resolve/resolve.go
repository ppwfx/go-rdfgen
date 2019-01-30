package resolve

import (
	"github.com/21stio/go-rdfgen/pkg/helper"
	"github.com/21stio/go-rdfgen/pkg/types"
	"github.com/21stio/go-rdfgen/pkg/parser"
	"github.com/21stio/go-rdfgen/pkg/utils"
	"strings"
)

func ResolveTriples(url string) (triples types.Triples, err error) {
	p := []types.Parser{parser.JsonLd{}, parser.Rdf{}}
	e := []types.IdExpander{helper.HashTag{}, helper.SchemaOrg{}, helper.PurlOrg{}, helper.Http{}, helper.DublinCoreOrg{}}

	knownUrls := map[string]bool{}
	queued := map[string]bool{}

	ch := make(chan string, 100000)

	ch <- url

FOR:
	for {
		select {
		case u := <-ch:
			if skip(u) {
				continue
			}

			_, ok := knownUrls[u]
			if ok {
				continue
			}
			knownUrls[u] = true

			var ts []types.Triple
			ts, err = getTriples(p, u)
			if err != nil {
				return
			}

			triples = append(triples, ts...)

			ids := getIds(ts)
			urls := idToUrls(e, ids)

			for _, u := range urls {
				_, ok := queued[u]
				if ok {
					continue
				}
				queued[u] = true

				ch <- u
			}
		default:
			break FOR
		}
	}

	return
}

func skip(u string) bool {
	if u == "http://www.w3.org/ns/regorg" {
		return true
	}

	if u == "http://pending.schema.org.jsonld" {
		return true
	}

	if u == "http://bib.schema.org.jsonld" {
		return true
	}

	if u == "http://meta.schema.org.jsonld" {
		return true
	}

	if u == "http://auto.schema.org.jsonld" {
		return true
	}

	if u == "http://health-lifesci.schema.org.jsonld" {
		return true
	}

	return false
}

func getIds(triples []types.Triple) (ids []string) {
	idsMap := map[string]bool{}

	for _, t := range triples {
		if strings.HasPrefix(t.Object, "http:") || strings.HasPrefix(t.Object, "https:") {
			idsMap[t.Object] = true
		}

		if strings.HasPrefix(t.Predicate, "http:") || strings.HasPrefix(t.Predicate, "https:") {
			idsMap[t.Predicate] = true
		}

		if strings.HasPrefix(t.Subject, "http:") || strings.HasPrefix(t.Subject, "https:") {
			idsMap[t.Subject] = true
		}
	}

	for id, _ := range idsMap {
		ids = append(ids, id)
	}

	return
}

func idToUrls(expanders []types.IdExpander, ids []string) (urls []string) {
	for _, id := range ids {
		for _, e := range expanders {
			if e.CanToUrl(id) {
				id = e.ToUrl(id)
			}
		}

		urls = append(urls, id)
	}

	return urls
}

func getTriples(parsers []types.Parser, url string) (triples types.Triples, err error) {
	r, err := utils.Get(url)
	if err != nil {
		return
	}

	for _, p := range parsers {
		if !p.CanParse(r) {
			continue
		}

		triples, err = p.Parse(r.Body)
		if err != nil {
			return
		}

		break
	}

	return
}
