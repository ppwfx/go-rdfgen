package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInfectiousDisease strings.Replacer
var strconvInfectiousDisease strconv.NumError

var basicInfectiousDiseaseTraitMapping = map[string]func(*schema.InfectiousDiseaseTrait, []string){}
var customInfectiousDiseaseTraitMapping = map[string]func(*schema.InfectiousDiseaseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InfectiousDisease", func(ctx jsonld.Context) (interface{}) {
		return NewInfectiousDiseaseFromContext(ctx)
	})


	basicInfectiousDiseaseTraitMapping["http://schema.org/transmissionMethod"] = func(x *schema.InfectiousDiseaseTrait, s []string) {
		x.TransmissionMethod = s[0]
	}


	basicInfectiousDiseaseTraitMapping["http://schema.org/infectiousAgent"] = func(x *schema.InfectiousDiseaseTrait, s []string) {
		x.InfectiousAgent = s[0]
	}



	customInfectiousDiseaseTraitMapping["http://schema.org/infectiousAgentClass"] = func(x *schema.InfectiousDiseaseTrait, ctx jsonld.Context, s string){
		var y schema.InfectiousAgentClass
		if strings.HasPrefix(s, "_:") {
			y = NewInfectiousAgentClassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInfectiousAgentClass()
			y.Id = s
		}

		x.InfectiousAgentClass = &y

		return
	}
}

func NewInfectiousDiseaseFromContext(ctx jsonld.Context) (x schema.InfectiousDisease) {
	x.Type = "http://schema.org/InfectiousDisease"
	MapBasicToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)
	MapCustomToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInfectiousDiseaseTrait(ctx, &x.InfectiousDiseaseTrait)
	MapCustomToInfectiousDiseaseTrait(ctx, &x.InfectiousDiseaseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInfectiousDiseaseTrait(ctx jsonld.Context, x *schema.InfectiousDiseaseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInfectiousDiseaseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInfectiousDiseaseTrait(ctx jsonld.Context, x *schema.InfectiousDiseaseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInfectiousDiseaseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}