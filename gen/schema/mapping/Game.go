package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGame strings.Replacer
var strconvGame strconv.NumError

var basicGameTraitMapping = map[string]func(*schema.GameTrait, []string){}
var customGameTraitMapping = map[string]func(*schema.GameTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Game", func(ctx jsonld.Context) (interface{}) {
		return NewGameFromContext(ctx)
	})







	customGameTraitMapping["http://schema.org/characterAttribute"] = func(x *schema.GameTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.CharacterAttribute = &y

		return
	}

	customGameTraitMapping["http://schema.org/gameItem"] = func(x *schema.GameTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.GameItem = &y

		return
	}

	customGameTraitMapping["http://schema.org/gameLocation"] = func(x *schema.GameTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GameLocation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GameLocation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GameLocation = s
		}
	}

	customGameTraitMapping["http://schema.org/numberOfPlayers"] = func(x *schema.GameTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.NumberOfPlayers = &y

		return
	}

	customGameTraitMapping["http://schema.org/quest"] = func(x *schema.GameTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Quest = &y

		return
	}
}

func NewGameFromContext(ctx jsonld.Context) (x schema.Game) {
	x.Type = "http://schema.org/Game"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGameTrait(ctx, &x.GameTrait)
	MapCustomToGameTrait(ctx, &x.GameTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGameTrait(ctx jsonld.Context, x *schema.GameTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGameTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGameTrait(ctx jsonld.Context, x *schema.GameTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGameTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}