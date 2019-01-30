package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLoseAction strings.Replacer
var strconvLoseAction strconv.NumError

var basicLoseActionTraitMapping = map[string]func(*schema.LoseActionTrait, []string){}
var customLoseActionTraitMapping = map[string]func(*schema.LoseActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LoseAction", func(ctx jsonld.Context) (interface{}) {
		return NewLoseActionFromContext(ctx)
	})



	customLoseActionTraitMapping["http://schema.org/winner"] = func(x *schema.LoseActionTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Winner = &y

		return
	}
}

func NewLoseActionFromContext(ctx jsonld.Context) (x schema.LoseAction) {
	x.Type = "http://schema.org/LoseAction"
	MapBasicToAchieveActionTrait(ctx, &x.AchieveActionTrait)
	MapCustomToAchieveActionTrait(ctx, &x.AchieveActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLoseActionTrait(ctx, &x.LoseActionTrait)
	MapCustomToLoseActionTrait(ctx, &x.LoseActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLoseActionTrait(ctx jsonld.Context, x *schema.LoseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLoseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLoseActionTrait(ctx jsonld.Context, x *schema.LoseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLoseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}