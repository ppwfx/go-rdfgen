package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWriteAction strings.Replacer
var strconvWriteAction strconv.NumError

var basicWriteActionTraitMapping = map[string]func(*schema.WriteActionTrait, []string){}
var customWriteActionTraitMapping = map[string]func(*schema.WriteActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WriteAction", func(ctx jsonld.Context) (interface{}) {
		return NewWriteActionFromContext(ctx)
	})




	customWriteActionTraitMapping["http://schema.org/inLanguage"] = func(x *schema.WriteActionTrait, ctx jsonld.Context, s string){
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

	customWriteActionTraitMapping["http://schema.org/language"] = func(x *schema.WriteActionTrait, ctx jsonld.Context, s string){
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
}

func NewWriteActionFromContext(ctx jsonld.Context) (x schema.WriteAction) {
	x.Type = "http://schema.org/WriteAction"
	MapBasicToCreateActionTrait(ctx, &x.CreateActionTrait)
	MapCustomToCreateActionTrait(ctx, &x.CreateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWriteActionTrait(ctx, &x.WriteActionTrait)
	MapCustomToWriteActionTrait(ctx, &x.WriteActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWriteActionTrait(ctx jsonld.Context, x *schema.WriteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWriteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWriteActionTrait(ctx jsonld.Context, x *schema.WriteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWriteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}