package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSteeringPositionValue strings.Replacer
var strconvSteeringPositionValue strconv.NumError

var basicSteeringPositionValueTraitMapping = map[string]func(*schema.SteeringPositionValueTrait, []string){}
var customSteeringPositionValueTraitMapping = map[string]func(*schema.SteeringPositionValueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SteeringPositionValue", func(ctx jsonld.Context) (interface{}) {
		return NewSteeringPositionValueFromContext(ctx)
	})

}

func NewSteeringPositionValueFromContext(ctx jsonld.Context) (x schema.SteeringPositionValue) {
	x.Type = "http://schema.org/SteeringPositionValue"
	MapBasicToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)
	MapCustomToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSteeringPositionValueTrait(ctx, &x.SteeringPositionValueTrait)
	MapCustomToSteeringPositionValueTrait(ctx, &x.SteeringPositionValueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSteeringPositionValueTrait(ctx jsonld.Context, x *schema.SteeringPositionValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSteeringPositionValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSteeringPositionValueTrait(ctx jsonld.Context, x *schema.SteeringPositionValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSteeringPositionValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}