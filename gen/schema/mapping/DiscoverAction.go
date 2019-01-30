package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDiscoverAction strings.Replacer
var strconvDiscoverAction strconv.NumError

var basicDiscoverActionTraitMapping = map[string]func(*schema.DiscoverActionTrait, []string){}
var customDiscoverActionTraitMapping = map[string]func(*schema.DiscoverActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DiscoverAction", func(ctx jsonld.Context) (interface{}) {
		return NewDiscoverActionFromContext(ctx)
	})

}

func NewDiscoverActionFromContext(ctx jsonld.Context) (x schema.DiscoverAction) {
	x.Type = "http://schema.org/DiscoverAction"
	MapBasicToFindActionTrait(ctx, &x.FindActionTrait)
	MapCustomToFindActionTrait(ctx, &x.FindActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDiscoverActionTrait(ctx, &x.DiscoverActionTrait)
	MapCustomToDiscoverActionTrait(ctx, &x.DiscoverActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDiscoverActionTrait(ctx jsonld.Context, x *schema.DiscoverActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDiscoverActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDiscoverActionTrait(ctx jsonld.Context, x *schema.DiscoverActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDiscoverActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}