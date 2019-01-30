package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTaxiStand strings.Replacer
var strconvTaxiStand strconv.NumError

var basicTaxiStandTraitMapping = map[string]func(*schema.TaxiStandTrait, []string){}
var customTaxiStandTraitMapping = map[string]func(*schema.TaxiStandTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TaxiStand", func(ctx jsonld.Context) (interface{}) {
		return NewTaxiStandFromContext(ctx)
	})

}

func NewTaxiStandFromContext(ctx jsonld.Context) (x schema.TaxiStand) {
	x.Type = "http://schema.org/TaxiStand"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTaxiStandTrait(ctx, &x.TaxiStandTrait)
	MapCustomToTaxiStandTrait(ctx, &x.TaxiStandTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTaxiStandTrait(ctx jsonld.Context, x *schema.TaxiStandTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTaxiStandTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTaxiStandTrait(ctx jsonld.Context, x *schema.TaxiStandTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTaxiStandTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}