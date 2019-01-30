package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalTrialDesign strings.Replacer
var strconvMedicalTrialDesign strconv.NumError

var basicMedicalTrialDesignTraitMapping = map[string]func(*schema.MedicalTrialDesignTrait, []string){}
var customMedicalTrialDesignTraitMapping = map[string]func(*schema.MedicalTrialDesignTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalTrialDesign", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalTrialDesignFromContext(ctx)
	})

}

func NewMedicalTrialDesignFromContext(ctx jsonld.Context) (x schema.MedicalTrialDesign) {
	x.Type = "http://schema.org/MedicalTrialDesign"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)


	MapBasicToMedicalTrialDesignTrait(ctx, &x.MedicalTrialDesignTrait)
	MapCustomToMedicalTrialDesignTrait(ctx, &x.MedicalTrialDesignTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalTrialDesignTrait(ctx jsonld.Context, x *schema.MedicalTrialDesignTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalTrialDesignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalTrialDesignTrait(ctx jsonld.Context, x *schema.MedicalTrialDesignTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalTrialDesignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}