package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBarOrPub strings.Replacer
var strconvBarOrPub strconv.NumError

var basicBarOrPubTraitMapping = map[string]func(*schema.BarOrPubTrait, []string){}
var customBarOrPubTraitMapping = map[string]func(*schema.BarOrPubTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BarOrPub", func(ctx jsonld.Context) (interface{}) {
		return NewBarOrPubFromContext(ctx)
	})

}

func NewBarOrPubFromContext(ctx jsonld.Context) (x schema.BarOrPub) {
	x.Type = "http://schema.org/BarOrPub"
	MapBasicToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)
	MapCustomToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToBarOrPubTrait(ctx, &x.BarOrPubTrait)
	MapCustomToBarOrPubTrait(ctx, &x.BarOrPubTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBarOrPubTrait(ctx jsonld.Context, x *schema.BarOrPubTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBarOrPubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBarOrPubTrait(ctx jsonld.Context, x *schema.BarOrPubTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBarOrPubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}