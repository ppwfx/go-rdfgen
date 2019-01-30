package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReactAction strings.Replacer
var strconvReactAction strconv.NumError

var basicReactActionTraitMapping = map[string]func(*schema.ReactActionTrait, []string){}
var customReactActionTraitMapping = map[string]func(*schema.ReactActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReactAction", func(ctx jsonld.Context) (interface{}) {
		return NewReactActionFromContext(ctx)
	})

}

func NewReactActionFromContext(ctx jsonld.Context) (x schema.ReactAction) {
	x.Type = "http://schema.org/ReactAction"
	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReactActionTrait(ctx, &x.ReactActionTrait)
	MapCustomToReactActionTrait(ctx, &x.ReactActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReactActionTrait(ctx jsonld.Context, x *schema.ReactActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReactActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReactActionTrait(ctx jsonld.Context, x *schema.ReactActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReactActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}