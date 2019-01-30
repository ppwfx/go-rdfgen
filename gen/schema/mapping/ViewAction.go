package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsViewAction strings.Replacer
var strconvViewAction strconv.NumError

var basicViewActionTraitMapping = map[string]func(*schema.ViewActionTrait, []string){}
var customViewActionTraitMapping = map[string]func(*schema.ViewActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ViewAction", func(ctx jsonld.Context) (interface{}) {
		return NewViewActionFromContext(ctx)
	})

}

func NewViewActionFromContext(ctx jsonld.Context) (x schema.ViewAction) {
	x.Type = "http://schema.org/ViewAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToViewActionTrait(ctx, &x.ViewActionTrait)
	MapCustomToViewActionTrait(ctx, &x.ViewActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToViewActionTrait(ctx jsonld.Context, x *schema.ViewActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicViewActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToViewActionTrait(ctx jsonld.Context, x *schema.ViewActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customViewActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}