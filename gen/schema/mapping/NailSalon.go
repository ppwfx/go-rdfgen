package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNailSalon strings.Replacer
var strconvNailSalon strconv.NumError

var basicNailSalonTraitMapping = map[string]func(*schema.NailSalonTrait, []string){}
var customNailSalonTraitMapping = map[string]func(*schema.NailSalonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/NailSalon", func(ctx jsonld.Context) (interface{}) {
		return NewNailSalonFromContext(ctx)
	})

}

func NewNailSalonFromContext(ctx jsonld.Context) (x schema.NailSalon) {
	x.Type = "http://schema.org/NailSalon"
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


	MapBasicToNailSalonTrait(ctx, &x.NailSalonTrait)
	MapCustomToNailSalonTrait(ctx, &x.NailSalonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNailSalonTrait(ctx jsonld.Context, x *schema.NailSalonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNailSalonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNailSalonTrait(ctx jsonld.Context, x *schema.NailSalonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNailSalonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}