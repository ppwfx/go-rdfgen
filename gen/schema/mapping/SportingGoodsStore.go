package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSportingGoodsStore strings.Replacer
var strconvSportingGoodsStore strconv.NumError

var basicSportingGoodsStoreTraitMapping = map[string]func(*schema.SportingGoodsStoreTrait, []string){}
var customSportingGoodsStoreTraitMapping = map[string]func(*schema.SportingGoodsStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SportingGoodsStore", func(ctx jsonld.Context) (interface{}) {
		return NewSportingGoodsStoreFromContext(ctx)
	})

}

func NewSportingGoodsStoreFromContext(ctx jsonld.Context) (x schema.SportingGoodsStore) {
	x.Type = "http://schema.org/SportingGoodsStore"
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


	MapBasicToSportingGoodsStoreTrait(ctx, &x.SportingGoodsStoreTrait)
	MapCustomToSportingGoodsStoreTrait(ctx, &x.SportingGoodsStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSportingGoodsStoreTrait(ctx jsonld.Context, x *schema.SportingGoodsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSportingGoodsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSportingGoodsStoreTrait(ctx jsonld.Context, x *schema.SportingGoodsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSportingGoodsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}