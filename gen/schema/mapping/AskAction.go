package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAskAction strings.Replacer
var strconvAskAction strconv.NumError

var basicAskActionTraitMapping = map[string]func(*schema.AskActionTrait, []string){}
var customAskActionTraitMapping = map[string]func(*schema.AskActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AskAction", func(ctx jsonld.Context) (interface{}) {
		return NewAskActionFromContext(ctx)
	})



	customAskActionTraitMapping["http://schema.org/question"] = func(x *schema.AskActionTrait, ctx jsonld.Context, s string){
		var y schema.Question
		if strings.HasPrefix(s, "_:") {
			y = NewQuestionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuestion()
			y.Id = s
		}

		x.Question = &y

		return
	}
}

func NewAskActionFromContext(ctx jsonld.Context) (x schema.AskAction) {
	x.Type = "http://schema.org/AskAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAskActionTrait(ctx, &x.AskActionTrait)
	MapCustomToAskActionTrait(ctx, &x.AskActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAskActionTrait(ctx jsonld.Context, x *schema.AskActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAskActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAskActionTrait(ctx jsonld.Context, x *schema.AskActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAskActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}