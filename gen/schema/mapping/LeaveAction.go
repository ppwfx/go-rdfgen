package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLeaveAction strings.Replacer
var strconvLeaveAction strconv.NumError

var basicLeaveActionTraitMapping = map[string]func(*schema.LeaveActionTrait, []string){}
var customLeaveActionTraitMapping = map[string]func(*schema.LeaveActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LeaveAction", func(ctx jsonld.Context) (interface{}) {
		return NewLeaveActionFromContext(ctx)
	})



	customLeaveActionTraitMapping["http://schema.org/event"] = func(x *schema.LeaveActionTrait, ctx jsonld.Context, s string){
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

func NewLeaveActionFromContext(ctx jsonld.Context) (x schema.LeaveAction) {
	x.Type = "http://schema.org/LeaveAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLeaveActionTrait(ctx, &x.LeaveActionTrait)
	MapCustomToLeaveActionTrait(ctx, &x.LeaveActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLeaveActionTrait(ctx jsonld.Context, x *schema.LeaveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLeaveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLeaveActionTrait(ctx jsonld.Context, x *schema.LeaveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLeaveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}