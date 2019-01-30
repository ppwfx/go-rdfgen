package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsJoinAction strings.Replacer
var strconvJoinAction strconv.NumError

var basicJoinActionTraitMapping = map[string]func(*schema.JoinActionTrait, []string){}
var customJoinActionTraitMapping = map[string]func(*schema.JoinActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/JoinAction", func(ctx jsonld.Context) (interface{}) {
		return NewJoinActionFromContext(ctx)
	})



	customJoinActionTraitMapping["http://schema.org/event"] = func(x *schema.JoinActionTrait, ctx jsonld.Context, s string){
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

func NewJoinActionFromContext(ctx jsonld.Context) (x schema.JoinAction) {
	x.Type = "http://schema.org/JoinAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToJoinActionTrait(ctx, &x.JoinActionTrait)
	MapCustomToJoinActionTrait(ctx, &x.JoinActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToJoinActionTrait(ctx jsonld.Context, x *schema.JoinActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicJoinActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToJoinActionTrait(ctx jsonld.Context, x *schema.JoinActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customJoinActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}