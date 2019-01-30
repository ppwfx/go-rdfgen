package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInteractionCounter strings.Replacer
var strconvInteractionCounter strconv.NumError

var basicInteractionCounterTraitMapping = map[string]func(*schema.InteractionCounterTrait, []string){}
var customInteractionCounterTraitMapping = map[string]func(*schema.InteractionCounterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InteractionCounter", func(ctx jsonld.Context) (interface{}) {
		return NewInteractionCounterFromContext(ctx)
	})





	customInteractionCounterTraitMapping["http://schema.org/interactionService"] = func(x *schema.InteractionCounterTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.InteractionService, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.InteractionService = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.InteractionService = s
		}
	}

	customInteractionCounterTraitMapping["http://schema.org/interactionType"] = func(x *schema.InteractionCounterTrait, ctx jsonld.Context, s string){
		var y schema.Action
		if strings.HasPrefix(s, "_:") {
			y = NewActionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAction()
			y.Id = s
		}

		x.InteractionType = &y

		return
	}

	customInteractionCounterTraitMapping["http://schema.org/userInteractionCount"] = func(x *schema.InteractionCounterTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.UserInteractionCount = &y

		return
	}
}

func NewInteractionCounterFromContext(ctx jsonld.Context) (x schema.InteractionCounter) {
	x.Type = "http://schema.org/InteractionCounter"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInteractionCounterTrait(ctx, &x.InteractionCounterTrait)
	MapCustomToInteractionCounterTrait(ctx, &x.InteractionCounterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInteractionCounterTrait(ctx jsonld.Context, x *schema.InteractionCounterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInteractionCounterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInteractionCounterTrait(ctx jsonld.Context, x *schema.InteractionCounterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInteractionCounterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}