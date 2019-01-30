package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDentist strings.Replacer
var strconvDentist strconv.NumError

var basicDentistTraitMapping = map[string]func(*schema.DentistTrait, []string){}
var customDentistTraitMapping = map[string]func(*schema.DentistTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Dentist", func(ctx jsonld.Context) (interface{}) {
		return NewDentistFromContext(ctx)
	})

}

func NewDentistFromContext(ctx jsonld.Context) (x schema.Dentist) {
	x.Type = "http://schema.org/Dentist"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)
	MapCustomToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)

	MapBasicToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)
	MapCustomToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)


	MapBasicToDentistTrait(ctx, &x.DentistTrait)
	MapCustomToDentistTrait(ctx, &x.DentistTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDentistTrait(ctx jsonld.Context, x *schema.DentistTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDentistTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDentistTrait(ctx jsonld.Context, x *schema.DentistTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDentistTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}