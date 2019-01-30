package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSocialEvent strings.Replacer
var strconvSocialEvent strconv.NumError

var basicSocialEventTraitMapping = map[string]func(*schema.SocialEventTrait, []string){}
var customSocialEventTraitMapping = map[string]func(*schema.SocialEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SocialEvent", func(ctx jsonld.Context) (interface{}) {
		return NewSocialEventFromContext(ctx)
	})

}

func NewSocialEventFromContext(ctx jsonld.Context) (x schema.SocialEvent) {
	x.Type = "http://schema.org/SocialEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSocialEventTrait(ctx, &x.SocialEventTrait)
	MapCustomToSocialEventTrait(ctx, &x.SocialEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSocialEventTrait(ctx jsonld.Context, x *schema.SocialEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSocialEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSocialEventTrait(ctx jsonld.Context, x *schema.SocialEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSocialEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}