package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsProfessionalService strings.Replacer
var strconvProfessionalService strconv.NumError

var basicProfessionalServiceTraitMapping = map[string]func(*schema.ProfessionalServiceTrait, []string){}
var customProfessionalServiceTraitMapping = map[string]func(*schema.ProfessionalServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ProfessionalService", func(ctx jsonld.Context) (interface{}) {
		return NewProfessionalServiceFromContext(ctx)
	})

}

func NewProfessionalServiceFromContext(ctx jsonld.Context) (x schema.ProfessionalService) {
	x.Type = "http://schema.org/ProfessionalService"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToProfessionalServiceTrait(ctx, &x.ProfessionalServiceTrait)
	MapCustomToProfessionalServiceTrait(ctx, &x.ProfessionalServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToProfessionalServiceTrait(ctx jsonld.Context, x *schema.ProfessionalServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicProfessionalServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToProfessionalServiceTrait(ctx jsonld.Context, x *schema.ProfessionalServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customProfessionalServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}