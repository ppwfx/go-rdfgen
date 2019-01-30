package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTattooParlor strings.Replacer
var strconvTattooParlor strconv.NumError

var basicTattooParlorTraitMapping = map[string]func(*schema.TattooParlorTrait, []string){}
var customTattooParlorTraitMapping = map[string]func(*schema.TattooParlorTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TattooParlor", func(ctx jsonld.Context) (interface{}) {
		return NewTattooParlorFromContext(ctx)
	})

}

func NewTattooParlorFromContext(ctx jsonld.Context) (x schema.TattooParlor) {
	x.Type = "http://schema.org/TattooParlor"
	MapBasicToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)
	MapCustomToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToTattooParlorTrait(ctx, &x.TattooParlorTrait)
	MapCustomToTattooParlorTrait(ctx, &x.TattooParlorTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTattooParlorTrait(ctx jsonld.Context, x *schema.TattooParlorTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTattooParlorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTattooParlorTrait(ctx jsonld.Context, x *schema.TattooParlorTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTattooParlorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}