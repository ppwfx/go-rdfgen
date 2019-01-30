package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInfectiousAgentClass strings.Replacer
var strconvInfectiousAgentClass strconv.NumError

var basicInfectiousAgentClassTraitMapping = map[string]func(*schema.InfectiousAgentClassTrait, []string){}
var customInfectiousAgentClassTraitMapping = map[string]func(*schema.InfectiousAgentClassTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InfectiousAgentClass", func(ctx jsonld.Context) (interface{}) {
		return NewInfectiousAgentClassFromContext(ctx)
	})

}

func NewInfectiousAgentClassFromContext(ctx jsonld.Context) (x schema.InfectiousAgentClass) {
	x.Type = "http://schema.org/InfectiousAgentClass"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInfectiousAgentClassTrait(ctx, &x.InfectiousAgentClassTrait)
	MapCustomToInfectiousAgentClassTrait(ctx, &x.InfectiousAgentClassTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInfectiousAgentClassTrait(ctx jsonld.Context, x *schema.InfectiousAgentClassTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInfectiousAgentClassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInfectiousAgentClassTrait(ctx jsonld.Context, x *schema.InfectiousAgentClassTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInfectiousAgentClassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}