package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMotel strings.Replacer
var strconvMotel strconv.NumError

var basicMotelTraitMapping = map[string]func(*schema.MotelTrait, []string){}
var customMotelTraitMapping = map[string]func(*schema.MotelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Motel", func(ctx jsonld.Context) (interface{}) {
		return NewMotelFromContext(ctx)
	})

}

func NewMotelFromContext(ctx jsonld.Context) (x schema.Motel) {
	x.Type = "http://schema.org/Motel"
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


	MapBasicToMotelTrait(ctx, &x.MotelTrait)
	MapCustomToMotelTrait(ctx, &x.MotelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMotelTrait(ctx jsonld.Context, x *schema.MotelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMotelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMotelTrait(ctx jsonld.Context, x *schema.MotelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMotelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}