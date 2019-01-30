package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInteractAction strings.Replacer
var strconvInteractAction strconv.NumError

var basicInteractActionTraitMapping = map[string]func(*schema.InteractActionTrait, []string){}
var customInteractActionTraitMapping = map[string]func(*schema.InteractActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InteractAction", func(ctx jsonld.Context) (interface{}) {
		return NewInteractActionFromContext(ctx)
	})

}

func NewInteractActionFromContext(ctx jsonld.Context) (x schema.InteractAction) {
	x.Type = "http://schema.org/InteractAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInteractActionTrait(ctx jsonld.Context, x *schema.InteractActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInteractActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInteractActionTrait(ctx jsonld.Context, x *schema.InteractActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInteractActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}