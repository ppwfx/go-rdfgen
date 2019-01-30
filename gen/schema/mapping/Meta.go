package mapping

import (
	"github.com/21stio/go-rdfgen/pkg/jsonld"
	"github.com/21stio/go-rdfgen/gen/schema"
)

var distinctMetaTraitMapping = map[string]func(*schema.MetaTrait, []string){}

func init() {
	distinctMetaTraitMapping["http://www.w3.org/1999/02/22-rdf-syntax-ns#type"] = func(x *schema.MetaTrait, s []string) {
		x.Type = s[0]
	}
}

func MapBasicToMetaTrait(ctx jsonld.Context, x *schema.MetaTrait) () {
	for k, v := range ctx.Current {
		f, ok := distinctMetaTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

