package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFlorist strings.Replacer
var strconvFlorist strconv.NumError

var basicFloristTraitMapping = map[string]func(*schema.FloristTrait, []string){}
var customFloristTraitMapping = map[string]func(*schema.FloristTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Florist", func(ctx jsonld.Context) (interface{}) {
		return NewFloristFromContext(ctx)
	})

}

func NewFloristFromContext(ctx jsonld.Context) (x schema.Florist) {
	x.Type = "http://schema.org/Florist"
	MapBasicToStoreTrait(ctx, &x.StoreTrait)
	MapCustomToStoreTrait(ctx, &x.StoreTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToFloristTrait(ctx, &x.FloristTrait)
	MapCustomToFloristTrait(ctx, &x.FloristTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFloristTrait(ctx jsonld.Context, x *schema.FloristTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFloristTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFloristTrait(ctx jsonld.Context, x *schema.FloristTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFloristTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}