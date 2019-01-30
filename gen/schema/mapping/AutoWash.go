package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutoWash strings.Replacer
var strconvAutoWash strconv.NumError

var basicAutoWashTraitMapping = map[string]func(*schema.AutoWashTrait, []string){}
var customAutoWashTraitMapping = map[string]func(*schema.AutoWashTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutoWash", func(ctx jsonld.Context) (interface{}) {
		return NewAutoWashFromContext(ctx)
	})

}

func NewAutoWashFromContext(ctx jsonld.Context) (x schema.AutoWash) {
	x.Type = "http://schema.org/AutoWash"
	MapBasicToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)
	MapCustomToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToAutoWashTrait(ctx, &x.AutoWashTrait)
	MapCustomToAutoWashTrait(ctx, &x.AutoWashTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutoWashTrait(ctx jsonld.Context, x *schema.AutoWashTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutoWashTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutoWashTrait(ctx jsonld.Context, x *schema.AutoWashTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutoWashTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}