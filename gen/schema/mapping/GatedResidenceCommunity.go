package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGatedResidenceCommunity strings.Replacer
var strconvGatedResidenceCommunity strconv.NumError

var basicGatedResidenceCommunityTraitMapping = map[string]func(*schema.GatedResidenceCommunityTrait, []string){}
var customGatedResidenceCommunityTraitMapping = map[string]func(*schema.GatedResidenceCommunityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GatedResidenceCommunity", func(ctx jsonld.Context) (interface{}) {
		return NewGatedResidenceCommunityFromContext(ctx)
	})

}

func NewGatedResidenceCommunityFromContext(ctx jsonld.Context) (x schema.GatedResidenceCommunity) {
	x.Type = "http://schema.org/GatedResidenceCommunity"
	MapBasicToResidenceTrait(ctx, &x.ResidenceTrait)
	MapCustomToResidenceTrait(ctx, &x.ResidenceTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGatedResidenceCommunityTrait(ctx, &x.GatedResidenceCommunityTrait)
	MapCustomToGatedResidenceCommunityTrait(ctx, &x.GatedResidenceCommunityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGatedResidenceCommunityTrait(ctx jsonld.Context, x *schema.GatedResidenceCommunityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGatedResidenceCommunityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGatedResidenceCommunityTrait(ctx jsonld.Context, x *schema.GatedResidenceCommunityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGatedResidenceCommunityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}