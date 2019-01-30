package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRadioChannel strings.Replacer
var strconvRadioChannel strconv.NumError

var basicRadioChannelTraitMapping = map[string]func(*schema.RadioChannelTrait, []string){}
var customRadioChannelTraitMapping = map[string]func(*schema.RadioChannelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RadioChannel", func(ctx jsonld.Context) (interface{}) {
		return NewRadioChannelFromContext(ctx)
	})

}

func NewRadioChannelFromContext(ctx jsonld.Context) (x schema.RadioChannel) {
	x.Type = "http://schema.org/RadioChannel"
	MapBasicToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)
	MapCustomToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRadioChannelTrait(ctx, &x.RadioChannelTrait)
	MapCustomToRadioChannelTrait(ctx, &x.RadioChannelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRadioChannelTrait(ctx jsonld.Context, x *schema.RadioChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRadioChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRadioChannelTrait(ctx jsonld.Context, x *schema.RadioChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRadioChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}