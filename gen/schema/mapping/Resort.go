package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsResort strings.Replacer
var strconvResort strconv.NumError

var basicResortTraitMapping = map[string]func(*schema.ResortTrait, []string){}
var customResortTraitMapping = map[string]func(*schema.ResortTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Resort", func(ctx jsonld.Context) (interface{}) {
		return NewResortFromContext(ctx)
	})

}

func NewResortFromContext(ctx jsonld.Context) (x schema.Resort) {
	x.Type = "http://schema.org/Resort"
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


	MapBasicToResortTrait(ctx, &x.ResortTrait)
	MapCustomToResortTrait(ctx, &x.ResortTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToResortTrait(ctx jsonld.Context, x *schema.ResortTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicResortTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToResortTrait(ctx jsonld.Context, x *schema.ResortTrait) () {
	for k, v := range ctx.Current {
		f, ok := customResortTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}