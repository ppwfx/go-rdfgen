package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPublicToilet strings.Replacer
var strconvPublicToilet strconv.NumError

var basicPublicToiletTraitMapping = map[string]func(*schema.PublicToiletTrait, []string){}
var customPublicToiletTraitMapping = map[string]func(*schema.PublicToiletTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PublicToilet", func(ctx jsonld.Context) (interface{}) {
		return NewPublicToiletFromContext(ctx)
	})

}

func NewPublicToiletFromContext(ctx jsonld.Context) (x schema.PublicToilet) {
	x.Type = "http://schema.org/PublicToilet"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPublicToiletTrait(ctx, &x.PublicToiletTrait)
	MapCustomToPublicToiletTrait(ctx, &x.PublicToiletTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPublicToiletTrait(ctx jsonld.Context, x *schema.PublicToiletTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPublicToiletTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPublicToiletTrait(ctx jsonld.Context, x *schema.PublicToiletTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPublicToiletTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}