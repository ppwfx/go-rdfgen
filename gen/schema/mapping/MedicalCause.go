package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalCause strings.Replacer
var strconvMedicalCause strconv.NumError

var basicMedicalCauseTraitMapping = map[string]func(*schema.MedicalCauseTrait, []string){}
var customMedicalCauseTraitMapping = map[string]func(*schema.MedicalCauseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalCause", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalCauseFromContext(ctx)
	})



	customMedicalCauseTraitMapping["http://schema.org/causeOf"] = func(x *schema.MedicalCauseTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.CauseOf = &y

		return
	}
}

func NewMedicalCauseFromContext(ctx jsonld.Context) (x schema.MedicalCause) {
	x.Type = "http://schema.org/MedicalCause"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalCauseTrait(ctx, &x.MedicalCauseTrait)
	MapCustomToMedicalCauseTrait(ctx, &x.MedicalCauseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalCauseTrait(ctx jsonld.Context, x *schema.MedicalCauseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalCauseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalCauseTrait(ctx jsonld.Context, x *schema.MedicalCauseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalCauseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}