package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsThesis strings.Replacer
var strconvThesis strconv.NumError

var basicThesisTraitMapping = map[string]func(*schema.ThesisTrait, []string){}
var customThesisTraitMapping = map[string]func(*schema.ThesisTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Thesis", func(ctx jsonld.Context) (interface{}) {
		return NewThesisFromContext(ctx)
	})


	basicThesisTraitMapping["http://schema.org/inSupportOf"] = func(x *schema.ThesisTrait, s []string) {
		x.InSupportOf = s[0]
	}

}

func NewThesisFromContext(ctx jsonld.Context) (x schema.Thesis) {
	x.Type = "http://schema.org/Thesis"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToThesisTrait(ctx, &x.ThesisTrait)
	MapCustomToThesisTrait(ctx, &x.ThesisTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToThesisTrait(ctx jsonld.Context, x *schema.ThesisTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicThesisTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToThesisTrait(ctx jsonld.Context, x *schema.ThesisTrait) () {
	for k, v := range ctx.Current {
		f, ok := customThesisTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}