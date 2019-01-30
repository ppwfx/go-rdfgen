package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHVACBusiness strings.Replacer
var strconvHVACBusiness strconv.NumError

var basicHVACBusinessTraitMapping = map[string]func(*schema.HVACBusinessTrait, []string){}
var customHVACBusinessTraitMapping = map[string]func(*schema.HVACBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HVACBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewHVACBusinessFromContext(ctx)
	})

}

func NewHVACBusinessFromContext(ctx jsonld.Context) (x schema.HVACBusiness) {
	x.Type = "http://schema.org/HVACBusiness"
	MapBasicToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)
	MapCustomToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToHVACBusinessTrait(ctx, &x.HVACBusinessTrait)
	MapCustomToHVACBusinessTrait(ctx, &x.HVACBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHVACBusinessTrait(ctx jsonld.Context, x *schema.HVACBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHVACBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHVACBusinessTrait(ctx jsonld.Context, x *schema.HVACBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHVACBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}