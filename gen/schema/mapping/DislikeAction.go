package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDislikeAction strings.Replacer
var strconvDislikeAction strconv.NumError

var basicDislikeActionTraitMapping = map[string]func(*schema.DislikeActionTrait, []string){}
var customDislikeActionTraitMapping = map[string]func(*schema.DislikeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DislikeAction", func(ctx jsonld.Context) (interface{}) {
		return NewDislikeActionFromContext(ctx)
	})

}

func NewDislikeActionFromContext(ctx jsonld.Context) (x schema.DislikeAction) {
	x.Type = "http://schema.org/DislikeAction"
	MapBasicToReactActionTrait(ctx, &x.ReactActionTrait)
	MapCustomToReactActionTrait(ctx, &x.ReactActionTrait)

	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDislikeActionTrait(ctx, &x.DislikeActionTrait)
	MapCustomToDislikeActionTrait(ctx, &x.DislikeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDislikeActionTrait(ctx jsonld.Context, x *schema.DislikeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDislikeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDislikeActionTrait(ctx jsonld.Context, x *schema.DislikeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDislikeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}