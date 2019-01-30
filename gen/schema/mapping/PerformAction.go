package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPerformAction strings.Replacer
var strconvPerformAction strconv.NumError

var basicPerformActionTraitMapping = map[string]func(*schema.PerformActionTrait, []string){}
var customPerformActionTraitMapping = map[string]func(*schema.PerformActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PerformAction", func(ctx jsonld.Context) (interface{}) {
		return NewPerformActionFromContext(ctx)
	})



	customPerformActionTraitMapping["http://schema.org/entertainmentBusiness"] = func(x *schema.PerformActionTrait, ctx jsonld.Context, s string){
		var y schema.EntertainmentBusiness
		if strings.HasPrefix(s, "_:") {
			y = NewEntertainmentBusinessFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEntertainmentBusiness()
			y.Id = s
		}

		x.EntertainmentBusiness = &y

		return
	}
}

func NewPerformActionFromContext(ctx jsonld.Context) (x schema.PerformAction) {
	x.Type = "http://schema.org/PerformAction"
	MapBasicToPlayActionTrait(ctx, &x.PlayActionTrait)
	MapCustomToPlayActionTrait(ctx, &x.PlayActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPerformActionTrait(ctx, &x.PerformActionTrait)
	MapCustomToPerformActionTrait(ctx, &x.PerformActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPerformActionTrait(ctx jsonld.Context, x *schema.PerformActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPerformActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPerformActionTrait(ctx jsonld.Context, x *schema.PerformActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPerformActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}