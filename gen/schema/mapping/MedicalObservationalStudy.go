package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalObservationalStudy strings.Replacer
var strconvMedicalObservationalStudy strconv.NumError

var basicMedicalObservationalStudyTraitMapping = map[string]func(*schema.MedicalObservationalStudyTrait, []string){}
var customMedicalObservationalStudyTraitMapping = map[string]func(*schema.MedicalObservationalStudyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalObservationalStudy", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalObservationalStudyFromContext(ctx)
	})



	customMedicalObservationalStudyTraitMapping["http://schema.org/studyDesign"] = func(x *schema.MedicalObservationalStudyTrait, ctx jsonld.Context, s string){
		var y schema.MedicalObservationalStudyDesign
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalObservationalStudyDesignFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalObservationalStudyDesign()
			y.Id = s
		}

		x.StudyDesign = &y

		return
	}
}

func NewMedicalObservationalStudyFromContext(ctx jsonld.Context) (x schema.MedicalObservationalStudy) {
	x.Type = "http://schema.org/MedicalObservationalStudy"
	MapBasicToMedicalStudyTrait(ctx, &x.MedicalStudyTrait)
	MapCustomToMedicalStudyTrait(ctx, &x.MedicalStudyTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalObservationalStudyTrait(ctx, &x.MedicalObservationalStudyTrait)
	MapCustomToMedicalObservationalStudyTrait(ctx, &x.MedicalObservationalStudyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalObservationalStudyTrait(ctx jsonld.Context, x *schema.MedicalObservationalStudyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalObservationalStudyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalObservationalStudyTrait(ctx jsonld.Context, x *schema.MedicalObservationalStudyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalObservationalStudyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}