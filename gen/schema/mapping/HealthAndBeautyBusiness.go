package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHealthAndBeautyBusiness strings.Replacer
var strconvHealthAndBeautyBusiness strconv.NumError

var basicHealthAndBeautyBusinessTraitMapping = map[string]func(*schema.HealthAndBeautyBusinessTrait, []string){}
var customHealthAndBeautyBusinessTraitMapping = map[string]func(*schema.HealthAndBeautyBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HealthAndBeautyBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewHealthAndBeautyBusinessFromContext(ctx)
	})

}

func NewHealthAndBeautyBusinessFromContext(ctx jsonld.Context) (x schema.HealthAndBeautyBusiness) {
	x.Type = "http://schema.org/HealthAndBeautyBusiness"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)
	MapCustomToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHealthAndBeautyBusinessTrait(ctx jsonld.Context, x *schema.HealthAndBeautyBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHealthAndBeautyBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHealthAndBeautyBusinessTrait(ctx jsonld.Context, x *schema.HealthAndBeautyBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHealthAndBeautyBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}