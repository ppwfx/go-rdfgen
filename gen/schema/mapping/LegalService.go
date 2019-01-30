package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLegalService strings.Replacer
var strconvLegalService strconv.NumError

var basicLegalServiceTraitMapping = map[string]func(*schema.LegalServiceTrait, []string){}
var customLegalServiceTraitMapping = map[string]func(*schema.LegalServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LegalService", func(ctx jsonld.Context) (interface{}) {
		return NewLegalServiceFromContext(ctx)
	})

}

func NewLegalServiceFromContext(ctx jsonld.Context) (x schema.LegalService) {
	x.Type = "http://schema.org/LegalService"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToLegalServiceTrait(ctx, &x.LegalServiceTrait)
	MapCustomToLegalServiceTrait(ctx, &x.LegalServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLegalServiceTrait(ctx jsonld.Context, x *schema.LegalServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLegalServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLegalServiceTrait(ctx jsonld.Context, x *schema.LegalServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLegalServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}