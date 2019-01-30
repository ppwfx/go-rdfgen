package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAdministrativeArea strings.Replacer
var strconvAdministrativeArea strconv.NumError

var basicAdministrativeAreaTraitMapping = map[string]func(*schema.AdministrativeAreaTrait, []string){}
var customAdministrativeAreaTraitMapping = map[string]func(*schema.AdministrativeAreaTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AdministrativeArea", func(ctx jsonld.Context) (interface{}) {
		return NewAdministrativeAreaFromContext(ctx)
	})

}

func NewAdministrativeAreaFromContext(ctx jsonld.Context) (x schema.AdministrativeArea) {
	x.Type = "http://schema.org/AdministrativeArea"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)
	MapCustomToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAdministrativeAreaTrait(ctx jsonld.Context, x *schema.AdministrativeAreaTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAdministrativeAreaTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAdministrativeAreaTrait(ctx jsonld.Context, x *schema.AdministrativeAreaTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAdministrativeAreaTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}