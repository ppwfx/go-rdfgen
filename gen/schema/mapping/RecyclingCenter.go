package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRecyclingCenter strings.Replacer
var strconvRecyclingCenter strconv.NumError

var basicRecyclingCenterTraitMapping = map[string]func(*schema.RecyclingCenterTrait, []string){}
var customRecyclingCenterTraitMapping = map[string]func(*schema.RecyclingCenterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RecyclingCenter", func(ctx jsonld.Context) (interface{}) {
		return NewRecyclingCenterFromContext(ctx)
	})

}

func NewRecyclingCenterFromContext(ctx jsonld.Context) (x schema.RecyclingCenter) {
	x.Type = "http://schema.org/RecyclingCenter"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToRecyclingCenterTrait(ctx, &x.RecyclingCenterTrait)
	MapCustomToRecyclingCenterTrait(ctx, &x.RecyclingCenterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRecyclingCenterTrait(ctx jsonld.Context, x *schema.RecyclingCenterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRecyclingCenterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRecyclingCenterTrait(ctx jsonld.Context, x *schema.RecyclingCenterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRecyclingCenterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}