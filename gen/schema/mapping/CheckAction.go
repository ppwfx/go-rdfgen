package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCheckAction strings.Replacer
var strconvCheckAction strconv.NumError

var basicCheckActionTraitMapping = map[string]func(*schema.CheckActionTrait, []string){}
var customCheckActionTraitMapping = map[string]func(*schema.CheckActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CheckAction", func(ctx jsonld.Context) (interface{}) {
		return NewCheckActionFromContext(ctx)
	})

}

func NewCheckActionFromContext(ctx jsonld.Context) (x schema.CheckAction) {
	x.Type = "http://schema.org/CheckAction"
	MapBasicToFindActionTrait(ctx, &x.FindActionTrait)
	MapCustomToFindActionTrait(ctx, &x.FindActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCheckActionTrait(ctx, &x.CheckActionTrait)
	MapCustomToCheckActionTrait(ctx, &x.CheckActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCheckActionTrait(ctx jsonld.Context, x *schema.CheckActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCheckActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCheckActionTrait(ctx jsonld.Context, x *schema.CheckActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCheckActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}