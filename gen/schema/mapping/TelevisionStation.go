package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTelevisionStation strings.Replacer
var strconvTelevisionStation strconv.NumError

var basicTelevisionStationTraitMapping = map[string]func(*schema.TelevisionStationTrait, []string){}
var customTelevisionStationTraitMapping = map[string]func(*schema.TelevisionStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TelevisionStation", func(ctx jsonld.Context) (interface{}) {
		return NewTelevisionStationFromContext(ctx)
	})

}

func NewTelevisionStationFromContext(ctx jsonld.Context) (x schema.TelevisionStation) {
	x.Type = "http://schema.org/TelevisionStation"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToTelevisionStationTrait(ctx, &x.TelevisionStationTrait)
	MapCustomToTelevisionStationTrait(ctx, &x.TelevisionStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTelevisionStationTrait(ctx jsonld.Context, x *schema.TelevisionStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTelevisionStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTelevisionStationTrait(ctx jsonld.Context, x *schema.TelevisionStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTelevisionStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}