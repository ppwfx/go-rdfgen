package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalAudience strings.Replacer
var strconvMedicalAudience strconv.NumError

var basicMedicalAudienceTraitMapping = map[string]func(*schema.MedicalAudienceTrait, []string){}
var customMedicalAudienceTraitMapping = map[string]func(*schema.MedicalAudienceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalAudience", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalAudienceFromContext(ctx)
	})

}

func NewMedicalAudienceFromContext(ctx jsonld.Context) (x schema.MedicalAudience) {
	x.Type = "http://schema.org/MedicalAudience"
	MapBasicToAudienceTrait(ctx, &x.AudienceTrait)
	MapCustomToAudienceTrait(ctx, &x.AudienceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)
	MapCustomToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)


	MapBasicToMedicalAudienceTrait(ctx, &x.MedicalAudienceTrait)
	MapCustomToMedicalAudienceTrait(ctx, &x.MedicalAudienceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalAudienceTrait(ctx jsonld.Context, x *schema.MedicalAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalAudienceTrait(ctx jsonld.Context, x *schema.MedicalAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}