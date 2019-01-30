package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMarryAction strings.Replacer
var strconvMarryAction strconv.NumError

var basicMarryActionTraitMapping = map[string]func(*schema.MarryActionTrait, []string){}
var customMarryActionTraitMapping = map[string]func(*schema.MarryActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MarryAction", func(ctx jsonld.Context) (interface{}) {
		return NewMarryActionFromContext(ctx)
	})

}

func NewMarryActionFromContext(ctx jsonld.Context) (x schema.MarryAction) {
	x.Type = "http://schema.org/MarryAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMarryActionTrait(ctx, &x.MarryActionTrait)
	MapCustomToMarryActionTrait(ctx, &x.MarryActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMarryActionTrait(ctx jsonld.Context, x *schema.MarryActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMarryActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMarryActionTrait(ctx jsonld.Context, x *schema.MarryActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMarryActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}