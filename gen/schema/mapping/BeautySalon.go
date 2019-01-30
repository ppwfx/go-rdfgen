package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBeautySalon strings.Replacer
var strconvBeautySalon strconv.NumError

var basicBeautySalonTraitMapping = map[string]func(*schema.BeautySalonTrait, []string){}
var customBeautySalonTraitMapping = map[string]func(*schema.BeautySalonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BeautySalon", func(ctx jsonld.Context) (interface{}) {
		return NewBeautySalonFromContext(ctx)
	})

}

func NewBeautySalonFromContext(ctx jsonld.Context) (x schema.BeautySalon) {
	x.Type = "http://schema.org/BeautySalon"
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


	MapBasicToBeautySalonTrait(ctx, &x.BeautySalonTrait)
	MapCustomToBeautySalonTrait(ctx, &x.BeautySalonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBeautySalonTrait(ctx jsonld.Context, x *schema.BeautySalonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBeautySalonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBeautySalonTrait(ctx jsonld.Context, x *schema.BeautySalonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBeautySalonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}