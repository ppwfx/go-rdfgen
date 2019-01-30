package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTennisComplex strings.Replacer
var strconvTennisComplex strconv.NumError

var basicTennisComplexTraitMapping = map[string]func(*schema.TennisComplexTrait, []string){}
var customTennisComplexTraitMapping = map[string]func(*schema.TennisComplexTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TennisComplex", func(ctx jsonld.Context) (interface{}) {
		return NewTennisComplexFromContext(ctx)
	})

}

func NewTennisComplexFromContext(ctx jsonld.Context) (x schema.TennisComplex) {
	x.Type = "http://schema.org/TennisComplex"
	MapBasicToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)
	MapCustomToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToTennisComplexTrait(ctx, &x.TennisComplexTrait)
	MapCustomToTennisComplexTrait(ctx, &x.TennisComplexTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTennisComplexTrait(ctx jsonld.Context, x *schema.TennisComplexTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTennisComplexTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTennisComplexTrait(ctx jsonld.Context, x *schema.TennisComplexTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTennisComplexTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}