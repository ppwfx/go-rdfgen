package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBridge strings.Replacer
var strconvBridge strconv.NumError

var basicBridgeTraitMapping = map[string]func(*schema.BridgeTrait, []string){}
var customBridgeTraitMapping = map[string]func(*schema.BridgeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Bridge", func(ctx jsonld.Context) (interface{}) {
		return NewBridgeFromContext(ctx)
	})

}

func NewBridgeFromContext(ctx jsonld.Context) (x schema.Bridge) {
	x.Type = "http://schema.org/Bridge"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBridgeTrait(ctx, &x.BridgeTrait)
	MapCustomToBridgeTrait(ctx, &x.BridgeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBridgeTrait(ctx jsonld.Context, x *schema.BridgeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBridgeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBridgeTrait(ctx jsonld.Context, x *schema.BridgeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBridgeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}