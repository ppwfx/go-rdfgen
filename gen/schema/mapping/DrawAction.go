package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrawAction strings.Replacer
var strconvDrawAction strconv.NumError

var basicDrawActionTraitMapping = map[string]func(*schema.DrawActionTrait, []string){}
var customDrawActionTraitMapping = map[string]func(*schema.DrawActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrawAction", func(ctx jsonld.Context) (interface{}) {
		return NewDrawActionFromContext(ctx)
	})

}

func NewDrawActionFromContext(ctx jsonld.Context) (x schema.DrawAction) {
	x.Type = "http://schema.org/DrawAction"
	MapBasicToCreateActionTrait(ctx, &x.CreateActionTrait)
	MapCustomToCreateActionTrait(ctx, &x.CreateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrawActionTrait(ctx, &x.DrawActionTrait)
	MapCustomToDrawActionTrait(ctx, &x.DrawActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrawActionTrait(ctx jsonld.Context, x *schema.DrawActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrawActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrawActionTrait(ctx jsonld.Context, x *schema.DrawActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrawActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}