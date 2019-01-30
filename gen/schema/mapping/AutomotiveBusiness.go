package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutomotiveBusiness strings.Replacer
var strconvAutomotiveBusiness strconv.NumError

var basicAutomotiveBusinessTraitMapping = map[string]func(*schema.AutomotiveBusinessTrait, []string){}
var customAutomotiveBusinessTraitMapping = map[string]func(*schema.AutomotiveBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutomotiveBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewAutomotiveBusinessFromContext(ctx)
	})

}

func NewAutomotiveBusinessFromContext(ctx jsonld.Context) (x schema.AutomotiveBusiness) {
	x.Type = "http://schema.org/AutomotiveBusiness"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)
	MapCustomToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutomotiveBusinessTrait(ctx jsonld.Context, x *schema.AutomotiveBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutomotiveBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutomotiveBusinessTrait(ctx jsonld.Context, x *schema.AutomotiveBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutomotiveBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}