package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPrependAction strings.Replacer
var strconvPrependAction strconv.NumError

var basicPrependActionTraitMapping = map[string]func(*schema.PrependActionTrait, []string){}
var customPrependActionTraitMapping = map[string]func(*schema.PrependActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PrependAction", func(ctx jsonld.Context) (interface{}) {
		return NewPrependActionFromContext(ctx)
	})

}

func NewPrependActionFromContext(ctx jsonld.Context) (x schema.PrependAction) {
	x.Type = "http://schema.org/PrependAction"
	MapBasicToInsertActionTrait(ctx, &x.InsertActionTrait)
	MapCustomToInsertActionTrait(ctx, &x.InsertActionTrait)

	MapBasicToAddActionTrait(ctx, &x.AddActionTrait)
	MapCustomToAddActionTrait(ctx, &x.AddActionTrait)

	MapBasicToUpdateActionTrait(ctx, &x.UpdateActionTrait)
	MapCustomToUpdateActionTrait(ctx, &x.UpdateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPrependActionTrait(ctx, &x.PrependActionTrait)
	MapCustomToPrependActionTrait(ctx, &x.PrependActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPrependActionTrait(ctx jsonld.Context, x *schema.PrependActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPrependActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPrependActionTrait(ctx jsonld.Context, x *schema.PrependActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPrependActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}