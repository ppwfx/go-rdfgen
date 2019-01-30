package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHomeAndConstructionBusiness strings.Replacer
var strconvHomeAndConstructionBusiness strconv.NumError

var basicHomeAndConstructionBusinessTraitMapping = map[string]func(*schema.HomeAndConstructionBusinessTrait, []string){}
var customHomeAndConstructionBusinessTraitMapping = map[string]func(*schema.HomeAndConstructionBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HomeAndConstructionBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewHomeAndConstructionBusinessFromContext(ctx)
	})

}

func NewHomeAndConstructionBusinessFromContext(ctx jsonld.Context) (x schema.HomeAndConstructionBusiness) {
	x.Type = "http://schema.org/HomeAndConstructionBusiness"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)
	MapCustomToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHomeAndConstructionBusinessTrait(ctx jsonld.Context, x *schema.HomeAndConstructionBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHomeAndConstructionBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHomeAndConstructionBusinessTrait(ctx jsonld.Context, x *schema.HomeAndConstructionBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHomeAndConstructionBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}