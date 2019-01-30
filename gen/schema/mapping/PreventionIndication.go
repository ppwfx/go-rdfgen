package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPreventionIndication strings.Replacer
var strconvPreventionIndication strconv.NumError

var basicPreventionIndicationTraitMapping = map[string]func(*schema.PreventionIndicationTrait, []string){}
var customPreventionIndicationTraitMapping = map[string]func(*schema.PreventionIndicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PreventionIndication", func(ctx jsonld.Context) (interface{}) {
		return NewPreventionIndicationFromContext(ctx)
	})

}

func NewPreventionIndicationFromContext(ctx jsonld.Context) (x schema.PreventionIndication) {
	x.Type = "http://schema.org/PreventionIndication"
	MapBasicToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)
	MapCustomToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPreventionIndicationTrait(ctx, &x.PreventionIndicationTrait)
	MapCustomToPreventionIndicationTrait(ctx, &x.PreventionIndicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPreventionIndicationTrait(ctx jsonld.Context, x *schema.PreventionIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPreventionIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPreventionIndicationTrait(ctx jsonld.Context, x *schema.PreventionIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPreventionIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}