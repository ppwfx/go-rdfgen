package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRealEstateAgent strings.Replacer
var strconvRealEstateAgent strconv.NumError

var basicRealEstateAgentTraitMapping = map[string]func(*schema.RealEstateAgentTrait, []string){}
var customRealEstateAgentTraitMapping = map[string]func(*schema.RealEstateAgentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RealEstateAgent", func(ctx jsonld.Context) (interface{}) {
		return NewRealEstateAgentFromContext(ctx)
	})

}

func NewRealEstateAgentFromContext(ctx jsonld.Context) (x schema.RealEstateAgent) {
	x.Type = "http://schema.org/RealEstateAgent"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToRealEstateAgentTrait(ctx, &x.RealEstateAgentTrait)
	MapCustomToRealEstateAgentTrait(ctx, &x.RealEstateAgentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRealEstateAgentTrait(ctx jsonld.Context, x *schema.RealEstateAgentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRealEstateAgentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRealEstateAgentTrait(ctx jsonld.Context, x *schema.RealEstateAgentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRealEstateAgentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}