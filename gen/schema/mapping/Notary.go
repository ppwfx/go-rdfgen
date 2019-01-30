package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNotary strings.Replacer
var strconvNotary strconv.NumError

var basicNotaryTraitMapping = map[string]func(*schema.NotaryTrait, []string){}
var customNotaryTraitMapping = map[string]func(*schema.NotaryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Notary", func(ctx jsonld.Context) (interface{}) {
		return NewNotaryFromContext(ctx)
	})

}

func NewNotaryFromContext(ctx jsonld.Context) (x schema.Notary) {
	x.Type = "http://schema.org/Notary"
	MapBasicToLegalServiceTrait(ctx, &x.LegalServiceTrait)
	MapCustomToLegalServiceTrait(ctx, &x.LegalServiceTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToNotaryTrait(ctx, &x.NotaryTrait)
	MapCustomToNotaryTrait(ctx, &x.NotaryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNotaryTrait(ctx jsonld.Context, x *schema.NotaryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNotaryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNotaryTrait(ctx jsonld.Context, x *schema.NotaryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNotaryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}