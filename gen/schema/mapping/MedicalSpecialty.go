package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalSpecialty strings.Replacer
var strconvMedicalSpecialty strconv.NumError

var basicMedicalSpecialtyTraitMapping = map[string]func(*schema.MedicalSpecialtyTrait, []string){}
var customMedicalSpecialtyTraitMapping = map[string]func(*schema.MedicalSpecialtyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalSpecialty", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalSpecialtyFromContext(ctx)
	})

}

func NewMedicalSpecialtyFromContext(ctx jsonld.Context) (x schema.MedicalSpecialty) {
	x.Type = "http://schema.org/MedicalSpecialty"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToSpecialtyTrait(ctx, &x.SpecialtyTrait)
	MapCustomToSpecialtyTrait(ctx, &x.SpecialtyTrait)


	MapBasicToMedicalSpecialtyTrait(ctx, &x.MedicalSpecialtyTrait)
	MapCustomToMedicalSpecialtyTrait(ctx, &x.MedicalSpecialtyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalSpecialtyTrait(ctx jsonld.Context, x *schema.MedicalSpecialtyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalSpecialtyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalSpecialtyTrait(ctx jsonld.Context, x *schema.MedicalSpecialtyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalSpecialtyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}