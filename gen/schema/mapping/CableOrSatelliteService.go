package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCableOrSatelliteService strings.Replacer
var strconvCableOrSatelliteService strconv.NumError

var basicCableOrSatelliteServiceTraitMapping = map[string]func(*schema.CableOrSatelliteServiceTrait, []string){}
var customCableOrSatelliteServiceTraitMapping = map[string]func(*schema.CableOrSatelliteServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CableOrSatelliteService", func(ctx jsonld.Context) (interface{}) {
		return NewCableOrSatelliteServiceFromContext(ctx)
	})

}

func NewCableOrSatelliteServiceFromContext(ctx jsonld.Context) (x schema.CableOrSatelliteService) {
	x.Type = "http://schema.org/CableOrSatelliteService"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCableOrSatelliteServiceTrait(ctx, &x.CableOrSatelliteServiceTrait)
	MapCustomToCableOrSatelliteServiceTrait(ctx, &x.CableOrSatelliteServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCableOrSatelliteServiceTrait(ctx jsonld.Context, x *schema.CableOrSatelliteServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCableOrSatelliteServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCableOrSatelliteServiceTrait(ctx jsonld.Context, x *schema.CableOrSatelliteServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCableOrSatelliteServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}