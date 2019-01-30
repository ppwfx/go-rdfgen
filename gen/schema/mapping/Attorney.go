package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAttorney strings.Replacer
var strconvAttorney strconv.NumError

var basicAttorneyTraitMapping = map[string]func(*schema.AttorneyTrait, []string){}
var customAttorneyTraitMapping = map[string]func(*schema.AttorneyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Attorney", func(ctx jsonld.Context) (interface{}) {
		return NewAttorneyFromContext(ctx)
	})

}

func NewAttorneyFromContext(ctx jsonld.Context) (x schema.Attorney) {
	x.Type = "http://schema.org/Attorney"
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


	MapBasicToAttorneyTrait(ctx, &x.AttorneyTrait)
	MapCustomToAttorneyTrait(ctx, &x.AttorneyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAttorneyTrait(ctx jsonld.Context, x *schema.AttorneyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAttorneyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAttorneyTrait(ctx jsonld.Context, x *schema.AttorneyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAttorneyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}