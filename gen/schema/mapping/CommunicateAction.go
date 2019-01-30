package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCommunicateAction strings.Replacer
var strconvCommunicateAction strconv.NumError

var basicCommunicateActionTraitMapping = map[string]func(*schema.CommunicateActionTrait, []string){}
var customCommunicateActionTraitMapping = map[string]func(*schema.CommunicateActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CommunicateAction", func(ctx jsonld.Context) (interface{}) {
		return NewCommunicateActionFromContext(ctx)
	})






	customCommunicateActionTraitMapping["http://schema.org/about"] = func(x *schema.CommunicateActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.About = &y

		return
	}

	customCommunicateActionTraitMapping["http://schema.org/inLanguage"] = func(x *schema.CommunicateActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.InLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.InLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.InLanguage = s
		}
	}

	customCommunicateActionTraitMapping["http://schema.org/language"] = func(x *schema.CommunicateActionTrait, ctx jsonld.Context, s string){
		var y schema.Language
		if strings.HasPrefix(s, "_:") {
			y = NewLanguageFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLanguage()
			y.Id = s
		}

		x.Language = &y

		return
	}

	customCommunicateActionTraitMapping["http://schema.org/recipient"] = func(x *schema.CommunicateActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Recipient, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Recipient = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Recipient = s
		}
	}
}

func NewCommunicateActionFromContext(ctx jsonld.Context) (x schema.CommunicateAction) {
	x.Type = "http://schema.org/CommunicateAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCommunicateActionTrait(ctx jsonld.Context, x *schema.CommunicateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCommunicateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCommunicateActionTrait(ctx jsonld.Context, x *schema.CommunicateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCommunicateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}