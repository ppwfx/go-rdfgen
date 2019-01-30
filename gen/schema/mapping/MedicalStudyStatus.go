package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalStudyStatus strings.Replacer
var strconvMedicalStudyStatus strconv.NumError

var basicMedicalStudyStatusTraitMapping = map[string]func(*schema.MedicalStudyStatusTrait, []string){}
var customMedicalStudyStatusTraitMapping = map[string]func(*schema.MedicalStudyStatusTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalStudyStatus", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalStudyStatusFromContext(ctx)
	})

}

func NewMedicalStudyStatusFromContext(ctx jsonld.Context) (x schema.MedicalStudyStatus) {
	x.Type = "http://schema.org/MedicalStudyStatus"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalStudyStatusTrait(ctx, &x.MedicalStudyStatusTrait)
	MapCustomToMedicalStudyStatusTrait(ctx, &x.MedicalStudyStatusTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalStudyStatusTrait(ctx jsonld.Context, x *schema.MedicalStudyStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalStudyStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalStudyStatusTrait(ctx jsonld.Context, x *schema.MedicalStudyStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalStudyStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}