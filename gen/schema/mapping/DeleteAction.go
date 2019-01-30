package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDeleteAction strings.Replacer
var strconvDeleteAction strconv.NumError

var basicDeleteActionTraitMapping = map[string]func(*schema.DeleteActionTrait, []string){}
var customDeleteActionTraitMapping = map[string]func(*schema.DeleteActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DeleteAction", func(ctx jsonld.Context) (interface{}) {
		return NewDeleteActionFromContext(ctx)
	})

}

func NewDeleteActionFromContext(ctx jsonld.Context) (x schema.DeleteAction) {
	x.Type = "http://schema.org/DeleteAction"
	MapBasicToUpdateActionTrait(ctx, &x.UpdateActionTrait)
	MapCustomToUpdateActionTrait(ctx, &x.UpdateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDeleteActionTrait(ctx, &x.DeleteActionTrait)
	MapCustomToDeleteActionTrait(ctx, &x.DeleteActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDeleteActionTrait(ctx jsonld.Context, x *schema.DeleteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDeleteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDeleteActionTrait(ctx jsonld.Context, x *schema.DeleteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDeleteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}