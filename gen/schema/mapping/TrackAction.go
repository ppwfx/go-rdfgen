package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTrackAction strings.Replacer
var strconvTrackAction strconv.NumError

var basicTrackActionTraitMapping = map[string]func(*schema.TrackActionTrait, []string){}
var customTrackActionTraitMapping = map[string]func(*schema.TrackActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TrackAction", func(ctx jsonld.Context) (interface{}) {
		return NewTrackActionFromContext(ctx)
	})



	customTrackActionTraitMapping["http://schema.org/deliveryMethod"] = func(x *schema.TrackActionTrait, ctx jsonld.Context, s string){
		var y schema.DeliveryMethod
		if strings.HasPrefix(s, "_:") {
			y = NewDeliveryMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDeliveryMethod()
			y.Id = s
		}

		x.DeliveryMethod = &y

		return
	}
}

func NewTrackActionFromContext(ctx jsonld.Context) (x schema.TrackAction) {
	x.Type = "http://schema.org/TrackAction"
	MapBasicToFindActionTrait(ctx, &x.FindActionTrait)
	MapCustomToFindActionTrait(ctx, &x.FindActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTrackActionTrait(ctx, &x.TrackActionTrait)
	MapCustomToTrackActionTrait(ctx, &x.TrackActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTrackActionTrait(ctx jsonld.Context, x *schema.TrackActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTrackActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTrackActionTrait(ctx jsonld.Context, x *schema.TrackActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTrackActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}