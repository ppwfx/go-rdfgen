package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDefenceEstablishment strings.Replacer
var strconvDefenceEstablishment strconv.NumError

var basicDefenceEstablishmentTraitMapping = map[string]func(*schema.DefenceEstablishmentTrait, []string){}
var customDefenceEstablishmentTraitMapping = map[string]func(*schema.DefenceEstablishmentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DefenceEstablishment", func(ctx jsonld.Context) (interface{}) {
		return NewDefenceEstablishmentFromContext(ctx)
	})

}

func NewDefenceEstablishmentFromContext(ctx jsonld.Context) (x schema.DefenceEstablishment) {
	x.Type = "http://schema.org/DefenceEstablishment"
	MapBasicToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)
	MapCustomToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDefenceEstablishmentTrait(ctx, &x.DefenceEstablishmentTrait)
	MapCustomToDefenceEstablishmentTrait(ctx, &x.DefenceEstablishmentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDefenceEstablishmentTrait(ctx jsonld.Context, x *schema.DefenceEstablishmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDefenceEstablishmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDefenceEstablishmentTrait(ctx jsonld.Context, x *schema.DefenceEstablishmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDefenceEstablishmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}