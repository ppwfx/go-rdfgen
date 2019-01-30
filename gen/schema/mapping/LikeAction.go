package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLikeAction strings.Replacer
var strconvLikeAction strconv.NumError

var basicLikeActionTraitMapping = map[string]func(*schema.LikeActionTrait, []string){}
var customLikeActionTraitMapping = map[string]func(*schema.LikeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LikeAction", func(ctx jsonld.Context) (interface{}) {
		return NewLikeActionFromContext(ctx)
	})

}

func NewLikeActionFromContext(ctx jsonld.Context) (x schema.LikeAction) {
	x.Type = "http://schema.org/LikeAction"
	MapBasicToReactActionTrait(ctx, &x.ReactActionTrait)
	MapCustomToReactActionTrait(ctx, &x.ReactActionTrait)

	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLikeActionTrait(ctx, &x.LikeActionTrait)
	MapCustomToLikeActionTrait(ctx, &x.LikeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLikeActionTrait(ctx jsonld.Context, x *schema.LikeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLikeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLikeActionTrait(ctx jsonld.Context, x *schema.LikeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLikeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}