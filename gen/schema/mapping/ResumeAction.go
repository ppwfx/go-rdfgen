package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsResumeAction strings.Replacer
var strconvResumeAction strconv.NumError

var basicResumeActionTraitMapping = map[string]func(*schema.ResumeActionTrait, []string){}
var customResumeActionTraitMapping = map[string]func(*schema.ResumeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ResumeAction", func(ctx jsonld.Context) (interface{}) {
		return NewResumeActionFromContext(ctx)
	})

}

func NewResumeActionFromContext(ctx jsonld.Context) (x schema.ResumeAction) {
	x.Type = "http://schema.org/ResumeAction"
	MapBasicToControlActionTrait(ctx, &x.ControlActionTrait)
	MapCustomToControlActionTrait(ctx, &x.ControlActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToResumeActionTrait(ctx, &x.ResumeActionTrait)
	MapCustomToResumeActionTrait(ctx, &x.ResumeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToResumeActionTrait(ctx jsonld.Context, x *schema.ResumeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicResumeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToResumeActionTrait(ctx jsonld.Context, x *schema.ResumeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customResumeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}