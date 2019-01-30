package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWantAction strings.Replacer
var strconvWantAction strconv.NumError

var basicWantActionTraitMapping = map[string]func(*schema.WantActionTrait, []string){}
var customWantActionTraitMapping = map[string]func(*schema.WantActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WantAction", func(ctx jsonld.Context) (interface{}) {
		return NewWantActionFromContext(ctx)
	})

}

func NewWantActionFromContext(ctx jsonld.Context) (x schema.WantAction) {
	x.Type = "http://schema.org/WantAction"
	MapBasicToReactActionTrait(ctx, &x.ReactActionTrait)
	MapCustomToReactActionTrait(ctx, &x.ReactActionTrait)

	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWantActionTrait(ctx, &x.WantActionTrait)
	MapCustomToWantActionTrait(ctx, &x.WantActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWantActionTrait(ctx jsonld.Context, x *schema.WantActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWantActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWantActionTrait(ctx jsonld.Context, x *schema.WantActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWantActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}