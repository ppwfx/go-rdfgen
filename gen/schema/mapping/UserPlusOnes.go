package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserPlusOnes strings.Replacer
var strconvUserPlusOnes strconv.NumError

var basicUserPlusOnesTraitMapping = map[string]func(*schema.UserPlusOnesTrait, []string){}
var customUserPlusOnesTraitMapping = map[string]func(*schema.UserPlusOnesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserPlusOnes", func(ctx jsonld.Context) (interface{}) {
		return NewUserPlusOnesFromContext(ctx)
	})

}

func NewUserPlusOnesFromContext(ctx jsonld.Context) (x schema.UserPlusOnes) {
	x.Type = "http://schema.org/UserPlusOnes"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserPlusOnesTrait(ctx, &x.UserPlusOnesTrait)
	MapCustomToUserPlusOnesTrait(ctx, &x.UserPlusOnesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserPlusOnesTrait(ctx jsonld.Context, x *schema.UserPlusOnesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserPlusOnesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserPlusOnesTrait(ctx jsonld.Context, x *schema.UserPlusOnesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserPlusOnesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}