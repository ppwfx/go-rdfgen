package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTravelAction strings.Replacer
var strconvTravelAction strconv.NumError

var basicTravelActionTraitMapping = map[string]func(*schema.TravelActionTrait, []string){}
var customTravelActionTraitMapping = map[string]func(*schema.TravelActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TravelAction", func(ctx jsonld.Context) (interface{}) {
		return NewTravelActionFromContext(ctx)
	})



	customTravelActionTraitMapping["http://schema.org/distance"] = func(x *schema.TravelActionTrait, ctx jsonld.Context, s string){
		var y schema.Distance
		if strings.HasPrefix(s, "_:") {
			y = NewDistanceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDistance()
			y.Id = s
		}

		x.Distance = &y

		return
	}
}

func NewTravelActionFromContext(ctx jsonld.Context) (x schema.TravelAction) {
	x.Type = "http://schema.org/TravelAction"
	MapBasicToMoveActionTrait(ctx, &x.MoveActionTrait)
	MapCustomToMoveActionTrait(ctx, &x.MoveActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTravelActionTrait(ctx, &x.TravelActionTrait)
	MapCustomToTravelActionTrait(ctx, &x.TravelActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTravelActionTrait(ctx jsonld.Context, x *schema.TravelActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTravelActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTravelActionTrait(ctx jsonld.Context, x *schema.TravelActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTravelActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}