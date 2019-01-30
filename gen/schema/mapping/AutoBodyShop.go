package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutoBodyShop strings.Replacer
var strconvAutoBodyShop strconv.NumError

var basicAutoBodyShopTraitMapping = map[string]func(*schema.AutoBodyShopTrait, []string){}
var customAutoBodyShopTraitMapping = map[string]func(*schema.AutoBodyShopTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutoBodyShop", func(ctx jsonld.Context) (interface{}) {
		return NewAutoBodyShopFromContext(ctx)
	})

}

func NewAutoBodyShopFromContext(ctx jsonld.Context) (x schema.AutoBodyShop) {
	x.Type = "http://schema.org/AutoBodyShop"
	MapBasicToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)
	MapCustomToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToAutoBodyShopTrait(ctx, &x.AutoBodyShopTrait)
	MapCustomToAutoBodyShopTrait(ctx, &x.AutoBodyShopTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutoBodyShopTrait(ctx jsonld.Context, x *schema.AutoBodyShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutoBodyShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutoBodyShopTrait(ctx jsonld.Context, x *schema.AutoBodyShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutoBodyShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}