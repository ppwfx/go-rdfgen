package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserInteraction strings.Replacer
var strconvUserInteraction strconv.NumError

var basicUserInteractionTraitMapping = map[string]func(*schema.UserInteractionTrait, []string){}
var customUserInteractionTraitMapping = map[string]func(*schema.UserInteractionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserInteraction", func(ctx jsonld.Context) (interface{}) {
		return NewUserInteractionFromContext(ctx)
	})

}

func NewUserInteractionFromContext(ctx jsonld.Context) (x schema.UserInteraction) {
	x.Type = "http://schema.org/UserInteraction"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserInteractionTrait(ctx jsonld.Context, x *schema.UserInteractionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserInteractionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserInteractionTrait(ctx jsonld.Context, x *schema.UserInteractionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserInteractionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}