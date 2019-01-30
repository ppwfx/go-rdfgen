package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTieAction strings.Replacer
var strconvTieAction strconv.NumError

var basicTieActionTraitMapping = map[string]func(*schema.TieActionTrait, []string){}
var customTieActionTraitMapping = map[string]func(*schema.TieActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TieAction", func(ctx jsonld.Context) (interface{}) {
		return NewTieActionFromContext(ctx)
	})

}

func NewTieActionFromContext(ctx jsonld.Context) (x schema.TieAction) {
	x.Type = "http://schema.org/TieAction"
	MapBasicToAchieveActionTrait(ctx, &x.AchieveActionTrait)
	MapCustomToAchieveActionTrait(ctx, &x.AchieveActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTieActionTrait(ctx, &x.TieActionTrait)
	MapCustomToTieActionTrait(ctx, &x.TieActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTieActionTrait(ctx jsonld.Context, x *schema.TieActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTieActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTieActionTrait(ctx jsonld.Context, x *schema.TieActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTieActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}