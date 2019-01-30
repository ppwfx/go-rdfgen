package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsChurch strings.Replacer
var strconvChurch strconv.NumError

var basicChurchTraitMapping = map[string]func(*schema.ChurchTrait, []string){}
var customChurchTraitMapping = map[string]func(*schema.ChurchTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Church", func(ctx jsonld.Context) (interface{}) {
		return NewChurchFromContext(ctx)
	})

}

func NewChurchFromContext(ctx jsonld.Context) (x schema.Church) {
	x.Type = "http://schema.org/Church"
	MapBasicToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)
	MapCustomToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToChurchTrait(ctx, &x.ChurchTrait)
	MapCustomToChurchTrait(ctx, &x.ChurchTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToChurchTrait(ctx jsonld.Context, x *schema.ChurchTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicChurchTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToChurchTrait(ctx jsonld.Context, x *schema.ChurchTrait) () {
	for k, v := range ctx.Current {
		f, ok := customChurchTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}