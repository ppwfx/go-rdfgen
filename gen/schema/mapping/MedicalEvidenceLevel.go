package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalEvidenceLevel strings.Replacer
var strconvMedicalEvidenceLevel strconv.NumError

var basicMedicalEvidenceLevelTraitMapping = map[string]func(*schema.MedicalEvidenceLevelTrait, []string){}
var customMedicalEvidenceLevelTraitMapping = map[string]func(*schema.MedicalEvidenceLevelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalEvidenceLevel", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalEvidenceLevelFromContext(ctx)
	})

}

func NewMedicalEvidenceLevelFromContext(ctx jsonld.Context) (x schema.MedicalEvidenceLevel) {
	x.Type = "http://schema.org/MedicalEvidenceLevel"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalEvidenceLevelTrait(ctx, &x.MedicalEvidenceLevelTrait)
	MapCustomToMedicalEvidenceLevelTrait(ctx, &x.MedicalEvidenceLevelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalEvidenceLevelTrait(ctx jsonld.Context, x *schema.MedicalEvidenceLevelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalEvidenceLevelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalEvidenceLevelTrait(ctx jsonld.Context, x *schema.MedicalEvidenceLevelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalEvidenceLevelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}