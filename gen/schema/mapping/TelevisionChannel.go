package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTelevisionChannel strings.Replacer
var strconvTelevisionChannel strconv.NumError

var basicTelevisionChannelTraitMapping = map[string]func(*schema.TelevisionChannelTrait, []string){}
var customTelevisionChannelTraitMapping = map[string]func(*schema.TelevisionChannelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TelevisionChannel", func(ctx jsonld.Context) (interface{}) {
		return NewTelevisionChannelFromContext(ctx)
	})

}

func NewTelevisionChannelFromContext(ctx jsonld.Context) (x schema.TelevisionChannel) {
	x.Type = "http://schema.org/TelevisionChannel"
	MapBasicToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)
	MapCustomToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTelevisionChannelTrait(ctx, &x.TelevisionChannelTrait)
	MapCustomToTelevisionChannelTrait(ctx, &x.TelevisionChannelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTelevisionChannelTrait(ctx jsonld.Context, x *schema.TelevisionChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTelevisionChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTelevisionChannelTrait(ctx jsonld.Context, x *schema.TelevisionChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTelevisionChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}