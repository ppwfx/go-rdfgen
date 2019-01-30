package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAssessAction strings.Replacer
var strconvAssessAction strconv.NumError

var basicAssessActionTraitMapping = map[string]func(*schema.AssessActionTrait, []string){}
var customAssessActionTraitMapping = map[string]func(*schema.AssessActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AssessAction", func(ctx jsonld.Context) (interface{}) {
		return NewAssessActionFromContext(ctx)
	})

}

func NewAssessActionFromContext(ctx jsonld.Context) (x schema.AssessAction) {
	x.Type = "http://schema.org/AssessAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAssessActionTrait(ctx jsonld.Context, x *schema.AssessActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAssessActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAssessActionTrait(ctx jsonld.Context, x *schema.AssessActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAssessActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}