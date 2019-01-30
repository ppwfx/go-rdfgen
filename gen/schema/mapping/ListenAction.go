package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsListenAction strings.Replacer
var strconvListenAction strconv.NumError

var basicListenActionTraitMapping = map[string]func(*schema.ListenActionTrait, []string){}
var customListenActionTraitMapping = map[string]func(*schema.ListenActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ListenAction", func(ctx jsonld.Context) (interface{}) {
		return NewListenActionFromContext(ctx)
	})

}

func NewListenActionFromContext(ctx jsonld.Context) (x schema.ListenAction) {
	x.Type = "http://schema.org/ListenAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToListenActionTrait(ctx, &x.ListenActionTrait)
	MapCustomToListenActionTrait(ctx, &x.ListenActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToListenActionTrait(ctx jsonld.Context, x *schema.ListenActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicListenActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToListenActionTrait(ctx jsonld.Context, x *schema.ListenActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customListenActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}