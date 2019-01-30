package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPhotographAction strings.Replacer
var strconvPhotographAction strconv.NumError

var basicPhotographActionTraitMapping = map[string]func(*schema.PhotographActionTrait, []string){}
var customPhotographActionTraitMapping = map[string]func(*schema.PhotographActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PhotographAction", func(ctx jsonld.Context) (interface{}) {
		return NewPhotographActionFromContext(ctx)
	})

}

func NewPhotographActionFromContext(ctx jsonld.Context) (x schema.PhotographAction) {
	x.Type = "http://schema.org/PhotographAction"
	MapBasicToCreateActionTrait(ctx, &x.CreateActionTrait)
	MapCustomToCreateActionTrait(ctx, &x.CreateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPhotographActionTrait(ctx, &x.PhotographActionTrait)
	MapCustomToPhotographActionTrait(ctx, &x.PhotographActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPhotographActionTrait(ctx jsonld.Context, x *schema.PhotographActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPhotographActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPhotographActionTrait(ctx jsonld.Context, x *schema.PhotographActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPhotographActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}