package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPerformingArtsTheater strings.Replacer
var strconvPerformingArtsTheater strconv.NumError

var basicPerformingArtsTheaterTraitMapping = map[string]func(*schema.PerformingArtsTheaterTrait, []string){}
var customPerformingArtsTheaterTraitMapping = map[string]func(*schema.PerformingArtsTheaterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PerformingArtsTheater", func(ctx jsonld.Context) (interface{}) {
		return NewPerformingArtsTheaterFromContext(ctx)
	})

}

func NewPerformingArtsTheaterFromContext(ctx jsonld.Context) (x schema.PerformingArtsTheater) {
	x.Type = "http://schema.org/PerformingArtsTheater"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPerformingArtsTheaterTrait(ctx, &x.PerformingArtsTheaterTrait)
	MapCustomToPerformingArtsTheaterTrait(ctx, &x.PerformingArtsTheaterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPerformingArtsTheaterTrait(ctx jsonld.Context, x *schema.PerformingArtsTheaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPerformingArtsTheaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPerformingArtsTheaterTrait(ctx jsonld.Context, x *schema.PerformingArtsTheaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPerformingArtsTheaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}