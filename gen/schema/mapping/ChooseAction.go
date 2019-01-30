package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsChooseAction strings.Replacer
var strconvChooseAction strconv.NumError

var basicChooseActionTraitMapping = map[string]func(*schema.ChooseActionTrait, []string){}
var customChooseActionTraitMapping = map[string]func(*schema.ChooseActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ChooseAction", func(ctx jsonld.Context) (interface{}) {
		return NewChooseActionFromContext(ctx)
	})




	customChooseActionTraitMapping["http://schema.org/actionOption"] = func(x *schema.ChooseActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ActionOption, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ActionOption = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ActionOption = s
		}
	}

	customChooseActionTraitMapping["http://schema.org/option"] = func(x *schema.ChooseActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Option, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Option = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Option = s
		}
	}
}

func NewChooseActionFromContext(ctx jsonld.Context) (x schema.ChooseAction) {
	x.Type = "http://schema.org/ChooseAction"
	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToChooseActionTrait(ctx, &x.ChooseActionTrait)
	MapCustomToChooseActionTrait(ctx, &x.ChooseActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToChooseActionTrait(ctx jsonld.Context, x *schema.ChooseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicChooseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToChooseActionTrait(ctx jsonld.Context, x *schema.ChooseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customChooseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}