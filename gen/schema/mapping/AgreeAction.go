package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAgreeAction strings.Replacer
var strconvAgreeAction strconv.NumError

var basicAgreeActionTraitMapping = map[string]func(*schema.AgreeActionTrait, []string){}
var customAgreeActionTraitMapping = map[string]func(*schema.AgreeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AgreeAction", func(ctx jsonld.Context) (interface{}) {
		return NewAgreeActionFromContext(ctx)
	})

}

func NewAgreeActionFromContext(ctx jsonld.Context) (x schema.AgreeAction) {
	x.Type = "http://schema.org/AgreeAction"
	MapBasicToReactActionTrait(ctx, &x.ReactActionTrait)
	MapCustomToReactActionTrait(ctx, &x.ReactActionTrait)

	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAgreeActionTrait(ctx, &x.AgreeActionTrait)
	MapCustomToAgreeActionTrait(ctx, &x.AgreeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAgreeActionTrait(ctx jsonld.Context, x *schema.AgreeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAgreeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAgreeActionTrait(ctx jsonld.Context, x *schema.AgreeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAgreeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}