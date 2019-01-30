package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWatchAction strings.Replacer
var strconvWatchAction strconv.NumError

var basicWatchActionTraitMapping = map[string]func(*schema.WatchActionTrait, []string){}
var customWatchActionTraitMapping = map[string]func(*schema.WatchActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WatchAction", func(ctx jsonld.Context) (interface{}) {
		return NewWatchActionFromContext(ctx)
	})

}

func NewWatchActionFromContext(ctx jsonld.Context) (x schema.WatchAction) {
	x.Type = "http://schema.org/WatchAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWatchActionTrait(ctx, &x.WatchActionTrait)
	MapCustomToWatchActionTrait(ctx, &x.WatchActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWatchActionTrait(ctx jsonld.Context, x *schema.WatchActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWatchActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWatchActionTrait(ctx jsonld.Context, x *schema.WatchActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWatchActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}