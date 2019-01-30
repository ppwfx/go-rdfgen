package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPlumber strings.Replacer
var strconvPlumber strconv.NumError

var basicPlumberTraitMapping = map[string]func(*schema.PlumberTrait, []string){}
var customPlumberTraitMapping = map[string]func(*schema.PlumberTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Plumber", func(ctx jsonld.Context) (interface{}) {
		return NewPlumberFromContext(ctx)
	})

}

func NewPlumberFromContext(ctx jsonld.Context) (x schema.Plumber) {
	x.Type = "http://schema.org/Plumber"
	MapBasicToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)
	MapCustomToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToPlumberTrait(ctx, &x.PlumberTrait)
	MapCustomToPlumberTrait(ctx, &x.PlumberTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPlumberTrait(ctx jsonld.Context, x *schema.PlumberTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPlumberTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPlumberTrait(ctx jsonld.Context, x *schema.PlumberTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPlumberTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}