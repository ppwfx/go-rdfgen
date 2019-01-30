package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEntertainmentBusiness strings.Replacer
var strconvEntertainmentBusiness strconv.NumError

var basicEntertainmentBusinessTraitMapping = map[string]func(*schema.EntertainmentBusinessTrait, []string){}
var customEntertainmentBusinessTraitMapping = map[string]func(*schema.EntertainmentBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EntertainmentBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewEntertainmentBusinessFromContext(ctx)
	})

}

func NewEntertainmentBusinessFromContext(ctx jsonld.Context) (x schema.EntertainmentBusiness) {
	x.Type = "http://schema.org/EntertainmentBusiness"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)
	MapCustomToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEntertainmentBusinessTrait(ctx jsonld.Context, x *schema.EntertainmentBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEntertainmentBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEntertainmentBusinessTrait(ctx jsonld.Context, x *schema.EntertainmentBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEntertainmentBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}