package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserCheckins strings.Replacer
var strconvUserCheckins strconv.NumError

var basicUserCheckinsTraitMapping = map[string]func(*schema.UserCheckinsTrait, []string){}
var customUserCheckinsTraitMapping = map[string]func(*schema.UserCheckinsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserCheckins", func(ctx jsonld.Context) (interface{}) {
		return NewUserCheckinsFromContext(ctx)
	})

}

func NewUserCheckinsFromContext(ctx jsonld.Context) (x schema.UserCheckins) {
	x.Type = "http://schema.org/UserCheckins"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserCheckinsTrait(ctx, &x.UserCheckinsTrait)
	MapCustomToUserCheckinsTrait(ctx, &x.UserCheckinsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserCheckinsTrait(ctx jsonld.Context, x *schema.UserCheckinsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserCheckinsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserCheckinsTrait(ctx jsonld.Context, x *schema.UserCheckinsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserCheckinsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}