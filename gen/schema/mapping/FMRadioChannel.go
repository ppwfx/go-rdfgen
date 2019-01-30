package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFMRadioChannel strings.Replacer
var strconvFMRadioChannel strconv.NumError

var basicFMRadioChannelTraitMapping = map[string]func(*schema.FMRadioChannelTrait, []string){}
var customFMRadioChannelTraitMapping = map[string]func(*schema.FMRadioChannelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FMRadioChannel", func(ctx jsonld.Context) (interface{}) {
		return NewFMRadioChannelFromContext(ctx)
	})

}

func NewFMRadioChannelFromContext(ctx jsonld.Context) (x schema.FMRadioChannel) {
	x.Type = "http://schema.org/FMRadioChannel"
	MapBasicToRadioChannelTrait(ctx, &x.RadioChannelTrait)
	MapCustomToRadioChannelTrait(ctx, &x.RadioChannelTrait)

	MapBasicToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)
	MapCustomToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFMRadioChannelTrait(ctx, &x.FMRadioChannelTrait)
	MapCustomToFMRadioChannelTrait(ctx, &x.FMRadioChannelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFMRadioChannelTrait(ctx jsonld.Context, x *schema.FMRadioChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFMRadioChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFMRadioChannelTrait(ctx jsonld.Context, x *schema.FMRadioChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFMRadioChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}