package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsExhibitionEvent strings.Replacer
var strconvExhibitionEvent strconv.NumError

var basicExhibitionEventTraitMapping = map[string]func(*schema.ExhibitionEventTrait, []string){}
var customExhibitionEventTraitMapping = map[string]func(*schema.ExhibitionEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ExhibitionEvent", func(ctx jsonld.Context) (interface{}) {
		return NewExhibitionEventFromContext(ctx)
	})

}

func NewExhibitionEventFromContext(ctx jsonld.Context) (x schema.ExhibitionEvent) {
	x.Type = "http://schema.org/ExhibitionEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToExhibitionEventTrait(ctx, &x.ExhibitionEventTrait)
	MapCustomToExhibitionEventTrait(ctx, &x.ExhibitionEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToExhibitionEventTrait(ctx jsonld.Context, x *schema.ExhibitionEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicExhibitionEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToExhibitionEventTrait(ctx jsonld.Context, x *schema.ExhibitionEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customExhibitionEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}