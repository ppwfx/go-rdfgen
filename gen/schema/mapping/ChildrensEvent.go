package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsChildrensEvent strings.Replacer
var strconvChildrensEvent strconv.NumError

var basicChildrensEventTraitMapping = map[string]func(*schema.ChildrensEventTrait, []string){}
var customChildrensEventTraitMapping = map[string]func(*schema.ChildrensEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ChildrensEvent", func(ctx jsonld.Context) (interface{}) {
		return NewChildrensEventFromContext(ctx)
	})

}

func NewChildrensEventFromContext(ctx jsonld.Context) (x schema.ChildrensEvent) {
	x.Type = "http://schema.org/ChildrensEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToChildrensEventTrait(ctx, &x.ChildrensEventTrait)
	MapCustomToChildrensEventTrait(ctx, &x.ChildrensEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToChildrensEventTrait(ctx jsonld.Context, x *schema.ChildrensEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicChildrensEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToChildrensEventTrait(ctx jsonld.Context, x *schema.ChildrensEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customChildrensEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}