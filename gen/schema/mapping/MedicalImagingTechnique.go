package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalImagingTechnique strings.Replacer
var strconvMedicalImagingTechnique strconv.NumError

var basicMedicalImagingTechniqueTraitMapping = map[string]func(*schema.MedicalImagingTechniqueTrait, []string){}
var customMedicalImagingTechniqueTraitMapping = map[string]func(*schema.MedicalImagingTechniqueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalImagingTechnique", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalImagingTechniqueFromContext(ctx)
	})

}

func NewMedicalImagingTechniqueFromContext(ctx jsonld.Context) (x schema.MedicalImagingTechnique) {
	x.Type = "http://schema.org/MedicalImagingTechnique"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalImagingTechniqueTrait(ctx, &x.MedicalImagingTechniqueTrait)
	MapCustomToMedicalImagingTechniqueTrait(ctx, &x.MedicalImagingTechniqueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalImagingTechniqueTrait(ctx jsonld.Context, x *schema.MedicalImagingTechniqueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalImagingTechniqueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalImagingTechniqueTrait(ctx jsonld.Context, x *schema.MedicalImagingTechniqueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalImagingTechniqueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}