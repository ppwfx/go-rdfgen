package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBefriendAction strings.Replacer
var strconvBefriendAction strconv.NumError

var basicBefriendActionTraitMapping = map[string]func(*schema.BefriendActionTrait, []string){}
var customBefriendActionTraitMapping = map[string]func(*schema.BefriendActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BefriendAction", func(ctx jsonld.Context) (interface{}) {
		return NewBefriendActionFromContext(ctx)
	})

}

func NewBefriendActionFromContext(ctx jsonld.Context) (x schema.BefriendAction) {
	x.Type = "http://schema.org/BefriendAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBefriendActionTrait(ctx, &x.BefriendActionTrait)
	MapCustomToBefriendActionTrait(ctx, &x.BefriendActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBefriendActionTrait(ctx jsonld.Context, x *schema.BefriendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBefriendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBefriendActionTrait(ctx jsonld.Context, x *schema.BefriendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBefriendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}