package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGameServerStatus strings.Replacer
var strconvGameServerStatus strconv.NumError

var basicGameServerStatusTraitMapping = map[string]func(*schema.GameServerStatusTrait, []string){}
var customGameServerStatusTraitMapping = map[string]func(*schema.GameServerStatusTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GameServerStatus", func(ctx jsonld.Context) (interface{}) {
		return NewGameServerStatusFromContext(ctx)
	})

}

func NewGameServerStatusFromContext(ctx jsonld.Context) (x schema.GameServerStatus) {
	x.Type = "http://schema.org/GameServerStatus"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGameServerStatusTrait(ctx, &x.GameServerStatusTrait)
	MapCustomToGameServerStatusTrait(ctx, &x.GameServerStatusTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGameServerStatusTrait(ctx jsonld.Context, x *schema.GameServerStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGameServerStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGameServerStatusTrait(ctx jsonld.Context, x *schema.GameServerStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGameServerStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}