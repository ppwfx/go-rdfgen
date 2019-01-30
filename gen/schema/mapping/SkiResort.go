package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSkiResort strings.Replacer
var strconvSkiResort strconv.NumError

var basicSkiResortTraitMapping = map[string]func(*schema.SkiResortTrait, []string){}
var customSkiResortTraitMapping = map[string]func(*schema.SkiResortTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SkiResort", func(ctx jsonld.Context) (interface{}) {
		return NewSkiResortFromContext(ctx)
	})

}

func NewSkiResortFromContext(ctx jsonld.Context) (x schema.SkiResort) {
	x.Type = "http://schema.org/SkiResort"
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


	MapBasicToSkiResortTrait(ctx, &x.SkiResortTrait)
	MapCustomToSkiResortTrait(ctx, &x.SkiResortTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSkiResortTrait(ctx jsonld.Context, x *schema.SkiResortTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSkiResortTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSkiResortTrait(ctx jsonld.Context, x *schema.SkiResortTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSkiResortTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}