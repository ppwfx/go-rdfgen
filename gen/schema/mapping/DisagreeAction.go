package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDisagreeAction strings.Replacer
var strconvDisagreeAction strconv.NumError

var basicDisagreeActionTraitMapping = map[string]func(*schema.DisagreeActionTrait, []string){}
var customDisagreeActionTraitMapping = map[string]func(*schema.DisagreeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DisagreeAction", func(ctx jsonld.Context) (interface{}) {
		return NewDisagreeActionFromContext(ctx)
	})

}

func NewDisagreeActionFromContext(ctx jsonld.Context) (x schema.DisagreeAction) {
	x.Type = "http://schema.org/DisagreeAction"
	MapBasicToReactActionTrait(ctx, &x.ReactActionTrait)
	MapCustomToReactActionTrait(ctx, &x.ReactActionTrait)

	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDisagreeActionTrait(ctx, &x.DisagreeActionTrait)
	MapCustomToDisagreeActionTrait(ctx, &x.DisagreeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDisagreeActionTrait(ctx jsonld.Context, x *schema.DisagreeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDisagreeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDisagreeActionTrait(ctx jsonld.Context, x *schema.DisagreeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDisagreeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}