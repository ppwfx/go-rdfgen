package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVisualArtsEvent strings.Replacer
var strconvVisualArtsEvent strconv.NumError

var basicVisualArtsEventTraitMapping = map[string]func(*schema.VisualArtsEventTrait, []string){}
var customVisualArtsEventTraitMapping = map[string]func(*schema.VisualArtsEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VisualArtsEvent", func(ctx jsonld.Context) (interface{}) {
		return NewVisualArtsEventFromContext(ctx)
	})

}

func NewVisualArtsEventFromContext(ctx jsonld.Context) (x schema.VisualArtsEvent) {
	x.Type = "http://schema.org/VisualArtsEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVisualArtsEventTrait(ctx, &x.VisualArtsEventTrait)
	MapCustomToVisualArtsEventTrait(ctx, &x.VisualArtsEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVisualArtsEventTrait(ctx jsonld.Context, x *schema.VisualArtsEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVisualArtsEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVisualArtsEventTrait(ctx jsonld.Context, x *schema.VisualArtsEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVisualArtsEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}