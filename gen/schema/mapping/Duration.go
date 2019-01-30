package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDuration strings.Replacer
var strconvDuration strconv.NumError

var basicDurationTraitMapping = map[string]func(*schema.DurationTrait, []string){}
var customDurationTraitMapping = map[string]func(*schema.DurationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Duration", func(ctx jsonld.Context) (interface{}) {
		return NewDurationFromContext(ctx)
	})

}

func NewDurationFromContext(ctx jsonld.Context) (x schema.Duration) {
	x.Type = "http://schema.org/Duration"
	MapBasicToQuantityTrait(ctx, &x.QuantityTrait)
	MapCustomToQuantityTrait(ctx, &x.QuantityTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDurationTrait(ctx, &x.DurationTrait)
	MapCustomToDurationTrait(ctx, &x.DurationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDurationTrait(ctx jsonld.Context, x *schema.DurationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDurationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDurationTrait(ctx jsonld.Context, x *schema.DurationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDurationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}