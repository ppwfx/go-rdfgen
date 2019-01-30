package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserBlocks strings.Replacer
var strconvUserBlocks strconv.NumError

var basicUserBlocksTraitMapping = map[string]func(*schema.UserBlocksTrait, []string){}
var customUserBlocksTraitMapping = map[string]func(*schema.UserBlocksTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserBlocks", func(ctx jsonld.Context) (interface{}) {
		return NewUserBlocksFromContext(ctx)
	})

}

func NewUserBlocksFromContext(ctx jsonld.Context) (x schema.UserBlocks) {
	x.Type = "http://schema.org/UserBlocks"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserBlocksTrait(ctx, &x.UserBlocksTrait)
	MapCustomToUserBlocksTrait(ctx, &x.UserBlocksTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserBlocksTrait(ctx jsonld.Context, x *schema.UserBlocksTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserBlocksTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserBlocksTrait(ctx jsonld.Context, x *schema.UserBlocksTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserBlocksTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}