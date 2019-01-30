package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCampground strings.Replacer
var strconvCampground strconv.NumError

var basicCampgroundTraitMapping = map[string]func(*schema.CampgroundTrait, []string){}
var customCampgroundTraitMapping = map[string]func(*schema.CampgroundTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Campground", func(ctx jsonld.Context) (interface{}) {
		return NewCampgroundFromContext(ctx)
	})

}

func NewCampgroundFromContext(ctx jsonld.Context) (x schema.Campground) {
	x.Type = "http://schema.org/Campground"
	MapBasicToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)
	MapCustomToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)


	MapBasicToCampgroundTrait(ctx, &x.CampgroundTrait)
	MapCustomToCampgroundTrait(ctx, &x.CampgroundTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCampgroundTrait(ctx jsonld.Context, x *schema.CampgroundTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCampgroundTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCampgroundTrait(ctx jsonld.Context, x *schema.CampgroundTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCampgroundTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}