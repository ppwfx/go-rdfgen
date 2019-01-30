package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsConsortium strings.Replacer
var strconvConsortium strconv.NumError

var basicConsortiumTraitMapping = map[string]func(*schema.ConsortiumTrait, []string){}
var customConsortiumTraitMapping = map[string]func(*schema.ConsortiumTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Consortium", func(ctx jsonld.Context) (interface{}) {
		return NewConsortiumFromContext(ctx)
	})

}

func NewConsortiumFromContext(ctx jsonld.Context) (x schema.Consortium) {
	x.Type = "http://schema.org/Consortium"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToConsortiumTrait(ctx, &x.ConsortiumTrait)
	MapCustomToConsortiumTrait(ctx, &x.ConsortiumTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToConsortiumTrait(ctx jsonld.Context, x *schema.ConsortiumTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicConsortiumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToConsortiumTrait(ctx jsonld.Context, x *schema.ConsortiumTrait) () {
	for k, v := range ctx.Current {
		f, ok := customConsortiumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}