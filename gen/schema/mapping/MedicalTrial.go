package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalTrial strings.Replacer
var strconvMedicalTrial strconv.NumError

var basicMedicalTrialTraitMapping = map[string]func(*schema.MedicalTrialTrait, []string){}
var customMedicalTrialTraitMapping = map[string]func(*schema.MedicalTrialTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalTrial", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalTrialFromContext(ctx)
	})


	basicMedicalTrialTraitMapping["http://schema.org/phase"] = func(x *schema.MedicalTrialTrait, s []string) {
		x.Phase = s[0]
	}



	customMedicalTrialTraitMapping["http://schema.org/trialDesign"] = func(x *schema.MedicalTrialTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTrialDesign
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTrialDesignFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTrialDesign()
			y.Id = s
		}

		x.TrialDesign = &y

		return
	}
}

func NewMedicalTrialFromContext(ctx jsonld.Context) (x schema.MedicalTrial) {
	x.Type = "http://schema.org/MedicalTrial"
	MapBasicToMedicalStudyTrait(ctx, &x.MedicalStudyTrait)
	MapCustomToMedicalStudyTrait(ctx, &x.MedicalStudyTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalTrialTrait(ctx, &x.MedicalTrialTrait)
	MapCustomToMedicalTrialTrait(ctx, &x.MedicalTrialTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalTrialTrait(ctx jsonld.Context, x *schema.MedicalTrialTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalTrialTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalTrialTrait(ctx jsonld.Context, x *schema.MedicalTrialTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalTrialTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}