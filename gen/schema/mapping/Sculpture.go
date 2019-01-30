package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSculpture strings.Replacer
var strconvSculpture strconv.NumError

var basicSculptureTraitMapping = map[string]func(*schema.SculptureTrait, []string){}
var customSculptureTraitMapping = map[string]func(*schema.SculptureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Sculpture", func(ctx jsonld.Context) (interface{}) {
		return NewSculptureFromContext(ctx)
	})

}

func NewSculptureFromContext(ctx jsonld.Context) (x schema.Sculpture) {
	x.Type = "http://schema.org/Sculpture"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSculptureTrait(ctx, &x.SculptureTrait)
	MapCustomToSculptureTrait(ctx, &x.SculptureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSculptureTrait(ctx jsonld.Context, x *schema.SculptureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSculptureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSculptureTrait(ctx jsonld.Context, x *schema.SculptureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSculptureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}