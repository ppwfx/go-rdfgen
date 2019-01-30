package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsIntangible strings.Replacer
var strconvIntangible strconv.NumError

var basicIntangibleTraitMapping = map[string]func(*schema.IntangibleTrait, []string){}
var customIntangibleTraitMapping = map[string]func(*schema.IntangibleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Intangible", func(ctx jsonld.Context) (interface{}) {
		return NewIntangibleFromContext(ctx)
	})

}

func NewIntangibleFromContext(ctx jsonld.Context) (x schema.Intangible) {
	x.Type = "http://schema.org/Intangible"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToIntangibleTrait(ctx jsonld.Context, x *schema.IntangibleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicIntangibleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToIntangibleTrait(ctx jsonld.Context, x *schema.IntangibleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customIntangibleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}