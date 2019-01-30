package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMuseum strings.Replacer
var strconvMuseum strconv.NumError

var basicMuseumTraitMapping = map[string]func(*schema.MuseumTrait, []string){}
var customMuseumTraitMapping = map[string]func(*schema.MuseumTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Museum", func(ctx jsonld.Context) (interface{}) {
		return NewMuseumFromContext(ctx)
	})

}

func NewMuseumFromContext(ctx jsonld.Context) (x schema.Museum) {
	x.Type = "http://schema.org/Museum"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMuseumTrait(ctx, &x.MuseumTrait)
	MapCustomToMuseumTrait(ctx, &x.MuseumTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMuseumTrait(ctx jsonld.Context, x *schema.MuseumTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMuseumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMuseumTrait(ctx jsonld.Context, x *schema.MuseumTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMuseumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}