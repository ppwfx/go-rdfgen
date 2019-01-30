package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMoveAction strings.Replacer
var strconvMoveAction strconv.NumError

var basicMoveActionTraitMapping = map[string]func(*schema.MoveActionTrait, []string){}
var customMoveActionTraitMapping = map[string]func(*schema.MoveActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MoveAction", func(ctx jsonld.Context) (interface{}) {
		return NewMoveActionFromContext(ctx)
	})




	customMoveActionTraitMapping["http://schema.org/fromLocation"] = func(x *schema.MoveActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.FromLocation = &y

		return
	}

	customMoveActionTraitMapping["http://schema.org/toLocation"] = func(x *schema.MoveActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ToLocation = &y

		return
	}
}

func NewMoveActionFromContext(ctx jsonld.Context) (x schema.MoveAction) {
	x.Type = "http://schema.org/MoveAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMoveActionTrait(ctx, &x.MoveActionTrait)
	MapCustomToMoveActionTrait(ctx, &x.MoveActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMoveActionTrait(ctx jsonld.Context, x *schema.MoveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMoveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMoveActionTrait(ctx jsonld.Context, x *schema.MoveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMoveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}