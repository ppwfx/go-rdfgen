package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalBusiness strings.Replacer
var strconvMedicalBusiness strconv.NumError

var basicMedicalBusinessTraitMapping = map[string]func(*schema.MedicalBusinessTrait, []string){}
var customMedicalBusinessTraitMapping = map[string]func(*schema.MedicalBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalBusinessFromContext(ctx)
	})

}

func NewMedicalBusinessFromContext(ctx jsonld.Context) (x schema.MedicalBusiness) {
	x.Type = "http://schema.org/MedicalBusiness"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)
	MapCustomToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalBusinessTrait(ctx jsonld.Context, x *schema.MedicalBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalBusinessTrait(ctx jsonld.Context, x *schema.MedicalBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}