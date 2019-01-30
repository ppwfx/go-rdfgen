package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAMRadioChannel strings.Replacer
var strconvAMRadioChannel strconv.NumError

var basicAMRadioChannelTraitMapping = map[string]func(*schema.AMRadioChannelTrait, []string){}
var customAMRadioChannelTraitMapping = map[string]func(*schema.AMRadioChannelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AMRadioChannel", func(ctx jsonld.Context) (interface{}) {
		return NewAMRadioChannelFromContext(ctx)
	})

}

func NewAMRadioChannelFromContext(ctx jsonld.Context) (x schema.AMRadioChannel) {
	x.Type = "http://schema.org/AMRadioChannel"
	MapBasicToRadioChannelTrait(ctx, &x.RadioChannelTrait)
	MapCustomToRadioChannelTrait(ctx, &x.RadioChannelTrait)

	MapBasicToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)
	MapCustomToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAMRadioChannelTrait(ctx, &x.AMRadioChannelTrait)
	MapCustomToAMRadioChannelTrait(ctx, &x.AMRadioChannelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAMRadioChannelTrait(ctx jsonld.Context, x *schema.AMRadioChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAMRadioChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAMRadioChannelTrait(ctx jsonld.Context, x *schema.AMRadioChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAMRadioChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}