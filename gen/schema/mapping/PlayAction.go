package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPlayAction strings.Replacer
var strconvPlayAction strconv.NumError

var basicPlayActionTraitMapping = map[string]func(*schema.PlayActionTrait, []string){}
var customPlayActionTraitMapping = map[string]func(*schema.PlayActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PlayAction", func(ctx jsonld.Context) (interface{}) {
		return NewPlayActionFromContext(ctx)
	})




	customPlayActionTraitMapping["http://schema.org/audience"] = func(x *schema.PlayActionTrait, ctx jsonld.Context, s string){
		var y schema.Audience
		if strings.HasPrefix(s, "_:") {
			y = NewAudienceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAudience()
			y.Id = s
		}

		x.Audience = &y

		return
	}

	customPlayActionTraitMapping["http://schema.org/event"] = func(x *schema.PlayActionTrait, ctx jsonld.Context, s string){
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

func NewPlayActionFromContext(ctx jsonld.Context) (x schema.PlayAction) {
	x.Type = "http://schema.org/PlayAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPlayActionTrait(ctx, &x.PlayActionTrait)
	MapCustomToPlayActionTrait(ctx, &x.PlayActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPlayActionTrait(ctx jsonld.Context, x *schema.PlayActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPlayActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPlayActionTrait(ctx jsonld.Context, x *schema.PlayActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPlayActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}