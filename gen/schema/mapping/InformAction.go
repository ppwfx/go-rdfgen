package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInformAction strings.Replacer
var strconvInformAction strconv.NumError

var basicInformActionTraitMapping = map[string]func(*schema.InformActionTrait, []string){}
var customInformActionTraitMapping = map[string]func(*schema.InformActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InformAction", func(ctx jsonld.Context) (interface{}) {
		return NewInformActionFromContext(ctx)
	})



	customInformActionTraitMapping["http://schema.org/event"] = func(x *schema.InformActionTrait, ctx jsonld.Context, s string){
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

func NewInformActionFromContext(ctx jsonld.Context) (x schema.InformAction) {
	x.Type = "http://schema.org/InformAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInformActionTrait(ctx, &x.InformActionTrait)
	MapCustomToInformActionTrait(ctx, &x.InformActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInformActionTrait(ctx jsonld.Context, x *schema.InformActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInformActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInformActionTrait(ctx jsonld.Context, x *schema.InformActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInformActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}