package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsURL strings.Replacer
var strconvURL strconv.NumError

var basicURLTraitMapping = map[string]func(*schema.URLTrait, []string){}
var customURLTraitMapping = map[string]func(*schema.URLTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/URL", func(ctx jsonld.Context) (interface{}) {
		return NewURLFromContext(ctx)
	})

}

func NewURLFromContext(ctx jsonld.Context) (x schema.URL) {
	x.Type = "http://schema.org/URL"
	MapBasicToTextTrait(ctx, &x.TextTrait)
	MapCustomToTextTrait(ctx, &x.TextTrait)


	MapBasicToURLTrait(ctx, &x.URLTrait)
	MapCustomToURLTrait(ctx, &x.URLTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToURLTrait(ctx jsonld.Context, x *schema.URLTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicURLTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToURLTrait(ctx jsonld.Context, x *schema.URLTrait) () {
	for k, v := range ctx.Current {
		f, ok := customURLTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}