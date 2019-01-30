package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRadioStation strings.Replacer
var strconvRadioStation strconv.NumError

var basicRadioStationTraitMapping = map[string]func(*schema.RadioStationTrait, []string){}
var customRadioStationTraitMapping = map[string]func(*schema.RadioStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RadioStation", func(ctx jsonld.Context) (interface{}) {
		return NewRadioStationFromContext(ctx)
	})

}

func NewRadioStationFromContext(ctx jsonld.Context) (x schema.RadioStation) {
	x.Type = "http://schema.org/RadioStation"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToRadioStationTrait(ctx, &x.RadioStationTrait)
	MapCustomToRadioStationTrait(ctx, &x.RadioStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRadioStationTrait(ctx jsonld.Context, x *schema.RadioStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRadioStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRadioStationTrait(ctx jsonld.Context, x *schema.RadioStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRadioStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}