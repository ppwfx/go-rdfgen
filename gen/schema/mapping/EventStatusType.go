package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEventStatusType strings.Replacer
var strconvEventStatusType strconv.NumError

var basicEventStatusTypeTraitMapping = map[string]func(*schema.EventStatusTypeTrait, []string){}
var customEventStatusTypeTraitMapping = map[string]func(*schema.EventStatusTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EventStatusType", func(ctx jsonld.Context) (interface{}) {
		return NewEventStatusTypeFromContext(ctx)
	})

}

func NewEventStatusTypeFromContext(ctx jsonld.Context) (x schema.EventStatusType) {
	x.Type = "http://schema.org/EventStatusType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEventStatusTypeTrait(ctx, &x.EventStatusTypeTrait)
	MapCustomToEventStatusTypeTrait(ctx, &x.EventStatusTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEventStatusTypeTrait(ctx jsonld.Context, x *schema.EventStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEventStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEventStatusTypeTrait(ctx jsonld.Context, x *schema.EventStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEventStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}