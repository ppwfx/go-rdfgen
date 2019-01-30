package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEducationEvent strings.Replacer
var strconvEducationEvent strconv.NumError

var basicEducationEventTraitMapping = map[string]func(*schema.EducationEventTrait, []string){}
var customEducationEventTraitMapping = map[string]func(*schema.EducationEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EducationEvent", func(ctx jsonld.Context) (interface{}) {
		return NewEducationEventFromContext(ctx)
	})

}

func NewEducationEventFromContext(ctx jsonld.Context) (x schema.EducationEvent) {
	x.Type = "http://schema.org/EducationEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEducationEventTrait(ctx, &x.EducationEventTrait)
	MapCustomToEducationEventTrait(ctx, &x.EducationEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEducationEventTrait(ctx jsonld.Context, x *schema.EducationEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEducationEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEducationEventTrait(ctx jsonld.Context, x *schema.EducationEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEducationEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}