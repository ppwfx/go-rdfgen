package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAchieveAction strings.Replacer
var strconvAchieveAction strconv.NumError

var basicAchieveActionTraitMapping = map[string]func(*schema.AchieveActionTrait, []string){}
var customAchieveActionTraitMapping = map[string]func(*schema.AchieveActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AchieveAction", func(ctx jsonld.Context) (interface{}) {
		return NewAchieveActionFromContext(ctx)
	})

}

func NewAchieveActionFromContext(ctx jsonld.Context) (x schema.AchieveAction) {
	x.Type = "http://schema.org/AchieveAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAchieveActionTrait(ctx, &x.AchieveActionTrait)
	MapCustomToAchieveActionTrait(ctx, &x.AchieveActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAchieveActionTrait(ctx jsonld.Context, x *schema.AchieveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAchieveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAchieveActionTrait(ctx jsonld.Context, x *schema.AchieveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAchieveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}