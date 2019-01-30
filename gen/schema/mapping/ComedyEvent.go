package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComedyEvent strings.Replacer
var strconvComedyEvent strconv.NumError

var basicComedyEventTraitMapping = map[string]func(*schema.ComedyEventTrait, []string){}
var customComedyEventTraitMapping = map[string]func(*schema.ComedyEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComedyEvent", func(ctx jsonld.Context) (interface{}) {
		return NewComedyEventFromContext(ctx)
	})

}

func NewComedyEventFromContext(ctx jsonld.Context) (x schema.ComedyEvent) {
	x.Type = "http://schema.org/ComedyEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToComedyEventTrait(ctx, &x.ComedyEventTrait)
	MapCustomToComedyEventTrait(ctx, &x.ComedyEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComedyEventTrait(ctx jsonld.Context, x *schema.ComedyEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComedyEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComedyEventTrait(ctx jsonld.Context, x *schema.ComedyEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComedyEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}