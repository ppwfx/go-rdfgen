package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTouristInformationCenter strings.Replacer
var strconvTouristInformationCenter strconv.NumError

var basicTouristInformationCenterTraitMapping = map[string]func(*schema.TouristInformationCenterTrait, []string){}
var customTouristInformationCenterTraitMapping = map[string]func(*schema.TouristInformationCenterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TouristInformationCenter", func(ctx jsonld.Context) (interface{}) {
		return NewTouristInformationCenterFromContext(ctx)
	})

}

func NewTouristInformationCenterFromContext(ctx jsonld.Context) (x schema.TouristInformationCenter) {
	x.Type = "http://schema.org/TouristInformationCenter"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToTouristInformationCenterTrait(ctx, &x.TouristInformationCenterTrait)
	MapCustomToTouristInformationCenterTrait(ctx, &x.TouristInformationCenterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTouristInformationCenterTrait(ctx jsonld.Context, x *schema.TouristInformationCenterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTouristInformationCenterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTouristInformationCenterTrait(ctx jsonld.Context, x *schema.TouristInformationCenterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTouristInformationCenterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}