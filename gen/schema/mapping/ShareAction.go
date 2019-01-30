package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsShareAction strings.Replacer
var strconvShareAction strconv.NumError

var basicShareActionTraitMapping = map[string]func(*schema.ShareActionTrait, []string){}
var customShareActionTraitMapping = map[string]func(*schema.ShareActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ShareAction", func(ctx jsonld.Context) (interface{}) {
		return NewShareActionFromContext(ctx)
	})

}

func NewShareActionFromContext(ctx jsonld.Context) (x schema.ShareAction) {
	x.Type = "http://schema.org/ShareAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToShareActionTrait(ctx, &x.ShareActionTrait)
	MapCustomToShareActionTrait(ctx, &x.ShareActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToShareActionTrait(ctx jsonld.Context, x *schema.ShareActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicShareActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToShareActionTrait(ctx jsonld.Context, x *schema.ShareActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customShareActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}