package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCheckInAction strings.Replacer
var strconvCheckInAction strconv.NumError

var basicCheckInActionTraitMapping = map[string]func(*schema.CheckInActionTrait, []string){}
var customCheckInActionTraitMapping = map[string]func(*schema.CheckInActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CheckInAction", func(ctx jsonld.Context) (interface{}) {
		return NewCheckInActionFromContext(ctx)
	})

}

func NewCheckInActionFromContext(ctx jsonld.Context) (x schema.CheckInAction) {
	x.Type = "http://schema.org/CheckInAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCheckInActionTrait(ctx, &x.CheckInActionTrait)
	MapCustomToCheckInActionTrait(ctx, &x.CheckInActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCheckInActionTrait(ctx jsonld.Context, x *schema.CheckInActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCheckInActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCheckInActionTrait(ctx jsonld.Context, x *schema.CheckInActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCheckInActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}