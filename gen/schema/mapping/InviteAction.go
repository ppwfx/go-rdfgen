package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInviteAction strings.Replacer
var strconvInviteAction strconv.NumError

var basicInviteActionTraitMapping = map[string]func(*schema.InviteActionTrait, []string){}
var customInviteActionTraitMapping = map[string]func(*schema.InviteActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InviteAction", func(ctx jsonld.Context) (interface{}) {
		return NewInviteActionFromContext(ctx)
	})



	customInviteActionTraitMapping["http://schema.org/event"] = func(x *schema.InviteActionTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.Event = &y

		return
	}
}

func NewInviteActionFromContext(ctx jsonld.Context) (x schema.InviteAction) {
	x.Type = "http://schema.org/InviteAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInviteActionTrait(ctx, &x.InviteActionTrait)
	MapCustomToInviteActionTrait(ctx, &x.InviteActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInviteActionTrait(ctx jsonld.Context, x *schema.InviteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInviteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInviteActionTrait(ctx jsonld.Context, x *schema.InviteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInviteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}