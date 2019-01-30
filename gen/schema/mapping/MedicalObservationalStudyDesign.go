package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalObservationalStudyDesign strings.Replacer
var strconvMedicalObservationalStudyDesign strconv.NumError

var basicMedicalObservationalStudyDesignTraitMapping = map[string]func(*schema.MedicalObservationalStudyDesignTrait, []string){}
var customMedicalObservationalStudyDesignTraitMapping = map[string]func(*schema.MedicalObservationalStudyDesignTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalObservationalStudyDesign", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalObservationalStudyDesignFromContext(ctx)
	})

}

func NewMedicalObservationalStudyDesignFromContext(ctx jsonld.Context) (x schema.MedicalObservationalStudyDesign) {
	x.Type = "http://schema.org/MedicalObservationalStudyDesign"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalObservationalStudyDesignTrait(ctx, &x.MedicalObservationalStudyDesignTrait)
	MapCustomToMedicalObservationalStudyDesignTrait(ctx, &x.MedicalObservationalStudyDesignTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalObservationalStudyDesignTrait(ctx jsonld.Context, x *schema.MedicalObservationalStudyDesignTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalObservationalStudyDesignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalObservationalStudyDesignTrait(ctx jsonld.Context, x *schema.MedicalObservationalStudyDesignTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalObservationalStudyDesignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}