package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHobbyShop strings.Replacer
var strconvHobbyShop strconv.NumError

var basicHobbyShopTraitMapping = map[string]func(*schema.HobbyShopTrait, []string){}
var customHobbyShopTraitMapping = map[string]func(*schema.HobbyShopTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HobbyShop", func(ctx jsonld.Context) (interface{}) {
		return NewHobbyShopFromContext(ctx)
	})

}

func NewHobbyShopFromContext(ctx jsonld.Context) (x schema.HobbyShop) {
	x.Type = "http://schema.org/HobbyShop"
	MapBasicToStoreTrait(ctx, &x.StoreTrait)
	MapCustomToStoreTrait(ctx, &x.StoreTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToHobbyShopTrait(ctx, &x.HobbyShopTrait)
	MapCustomToHobbyShopTrait(ctx, &x.HobbyShopTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHobbyShopTrait(ctx jsonld.Context, x *schema.HobbyShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHobbyShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHobbyShopTrait(ctx jsonld.Context, x *schema.HobbyShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHobbyShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}