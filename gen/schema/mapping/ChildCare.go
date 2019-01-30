package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsChildCare strings.Replacer
var strconvChildCare strconv.NumError

var basicChildCareTraitMapping = map[string]func(*schema.ChildCareTrait, []string){}
var customChildCareTraitMapping = map[string]func(*schema.ChildCareTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ChildCare", func(ctx jsonld.Context) (interface{}) {
		return NewChildCareFromContext(ctx)
	})

}

func NewChildCareFromContext(ctx jsonld.Context) (x schema.ChildCare) {
	x.Type = "http://schema.org/ChildCare"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToChildCareTrait(ctx, &x.ChildCareTrait)
	MapCustomToChildCareTrait(ctx, &x.ChildCareTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToChildCareTrait(ctx jsonld.Context, x *schema.ChildCareTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicChildCareTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToChildCareTrait(ctx jsonld.Context, x *schema.ChildCareTrait) () {
	for k, v := range ctx.Current {
		f, ok := customChildCareTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}