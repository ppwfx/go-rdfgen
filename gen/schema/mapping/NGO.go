package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNGO strings.Replacer
var strconvNGO strconv.NumError

var basicNGOTraitMapping = map[string]func(*schema.NGOTrait, []string){}
var customNGOTraitMapping = map[string]func(*schema.NGOTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/NGO", func(ctx jsonld.Context) (interface{}) {
		return NewNGOFromContext(ctx)
	})

}

func NewNGOFromContext(ctx jsonld.Context) (x schema.NGO) {
	x.Type = "http://schema.org/NGO"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToNGOTrait(ctx, &x.NGOTrait)
	MapCustomToNGOTrait(ctx, &x.NGOTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNGOTrait(ctx jsonld.Context, x *schema.NGOTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNGOTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNGOTrait(ctx jsonld.Context, x *schema.NGOTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNGOTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}