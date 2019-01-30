package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWinAction strings.Replacer
var strconvWinAction strconv.NumError

var basicWinActionTraitMapping = map[string]func(*schema.WinActionTrait, []string){}
var customWinActionTraitMapping = map[string]func(*schema.WinActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WinAction", func(ctx jsonld.Context) (interface{}) {
		return NewWinActionFromContext(ctx)
	})



	customWinActionTraitMapping["http://schema.org/loser"] = func(x *schema.WinActionTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Loser = &y

		return
	}
}

func NewWinActionFromContext(ctx jsonld.Context) (x schema.WinAction) {
	x.Type = "http://schema.org/WinAction"
	MapBasicToAchieveActionTrait(ctx, &x.AchieveActionTrait)
	MapCustomToAchieveActionTrait(ctx, &x.AchieveActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWinActionTrait(ctx, &x.WinActionTrait)
	MapCustomToWinActionTrait(ctx, &x.WinActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWinActionTrait(ctx jsonld.Context, x *schema.WinActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWinActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWinActionTrait(ctx jsonld.Context, x *schema.WinActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWinActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}