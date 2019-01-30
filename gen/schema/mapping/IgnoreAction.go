package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsIgnoreAction strings.Replacer
var strconvIgnoreAction strconv.NumError

var basicIgnoreActionTraitMapping = map[string]func(*schema.IgnoreActionTrait, []string){}
var customIgnoreActionTraitMapping = map[string]func(*schema.IgnoreActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/IgnoreAction", func(ctx jsonld.Context) (interface{}) {
		return NewIgnoreActionFromContext(ctx)
	})

}

func NewIgnoreActionFromContext(ctx jsonld.Context) (x schema.IgnoreAction) {
	x.Type = "http://schema.org/IgnoreAction"
	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToIgnoreActionTrait(ctx, &x.IgnoreActionTrait)
	MapCustomToIgnoreActionTrait(ctx, &x.IgnoreActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToIgnoreActionTrait(ctx jsonld.Context, x *schema.IgnoreActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicIgnoreActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToIgnoreActionTrait(ctx jsonld.Context, x *schema.IgnoreActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customIgnoreActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}