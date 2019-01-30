package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBowlingAlley strings.Replacer
var strconvBowlingAlley strconv.NumError

var basicBowlingAlleyTraitMapping = map[string]func(*schema.BowlingAlleyTrait, []string){}
var customBowlingAlleyTraitMapping = map[string]func(*schema.BowlingAlleyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BowlingAlley", func(ctx jsonld.Context) (interface{}) {
		return NewBowlingAlleyFromContext(ctx)
	})

}

func NewBowlingAlleyFromContext(ctx jsonld.Context) (x schema.BowlingAlley) {
	x.Type = "http://schema.org/BowlingAlley"
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


	MapBasicToBowlingAlleyTrait(ctx, &x.BowlingAlleyTrait)
	MapCustomToBowlingAlleyTrait(ctx, &x.BowlingAlleyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBowlingAlleyTrait(ctx jsonld.Context, x *schema.BowlingAlleyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBowlingAlleyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBowlingAlleyTrait(ctx jsonld.Context, x *schema.BowlingAlleyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBowlingAlleyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}