package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLiteraryEvent strings.Replacer
var strconvLiteraryEvent strconv.NumError

var basicLiteraryEventTraitMapping = map[string]func(*schema.LiteraryEventTrait, []string){}
var customLiteraryEventTraitMapping = map[string]func(*schema.LiteraryEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LiteraryEvent", func(ctx jsonld.Context) (interface{}) {
		return NewLiteraryEventFromContext(ctx)
	})

}

func NewLiteraryEventFromContext(ctx jsonld.Context) (x schema.LiteraryEvent) {
	x.Type = "http://schema.org/LiteraryEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLiteraryEventTrait(ctx, &x.LiteraryEventTrait)
	MapCustomToLiteraryEventTrait(ctx, &x.LiteraryEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLiteraryEventTrait(ctx jsonld.Context, x *schema.LiteraryEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLiteraryEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLiteraryEventTrait(ctx jsonld.Context, x *schema.LiteraryEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLiteraryEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}