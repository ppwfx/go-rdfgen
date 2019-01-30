package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserPlays strings.Replacer
var strconvUserPlays strconv.NumError

var basicUserPlaysTraitMapping = map[string]func(*schema.UserPlaysTrait, []string){}
var customUserPlaysTraitMapping = map[string]func(*schema.UserPlaysTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserPlays", func(ctx jsonld.Context) (interface{}) {
		return NewUserPlaysFromContext(ctx)
	})

}

func NewUserPlaysFromContext(ctx jsonld.Context) (x schema.UserPlays) {
	x.Type = "http://schema.org/UserPlays"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserPlaysTrait(ctx, &x.UserPlaysTrait)
	MapCustomToUserPlaysTrait(ctx, &x.UserPlaysTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserPlaysTrait(ctx jsonld.Context, x *schema.UserPlaysTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserPlaysTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserPlaysTrait(ctx jsonld.Context, x *schema.UserPlaysTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserPlaysTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}