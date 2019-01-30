package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusinessEvent strings.Replacer
var strconvBusinessEvent strconv.NumError

var basicBusinessEventTraitMapping = map[string]func(*schema.BusinessEventTrait, []string){}
var customBusinessEventTraitMapping = map[string]func(*schema.BusinessEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusinessEvent", func(ctx jsonld.Context) (interface{}) {
		return NewBusinessEventFromContext(ctx)
	})

}

func NewBusinessEventFromContext(ctx jsonld.Context) (x schema.BusinessEvent) {
	x.Type = "http://schema.org/BusinessEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusinessEventTrait(ctx, &x.BusinessEventTrait)
	MapCustomToBusinessEventTrait(ctx, &x.BusinessEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusinessEventTrait(ctx jsonld.Context, x *schema.BusinessEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusinessEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusinessEventTrait(ctx jsonld.Context, x *schema.BusinessEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusinessEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}