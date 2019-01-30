package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEndorseAction strings.Replacer
var strconvEndorseAction strconv.NumError

var basicEndorseActionTraitMapping = map[string]func(*schema.EndorseActionTrait, []string){}
var customEndorseActionTraitMapping = map[string]func(*schema.EndorseActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EndorseAction", func(ctx jsonld.Context) (interface{}) {
		return NewEndorseActionFromContext(ctx)
	})



	customEndorseActionTraitMapping["http://schema.org/endorsee"] = func(x *schema.EndorseActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Endorsee, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Endorsee = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Endorsee = s
		}
	}
}

func NewEndorseActionFromContext(ctx jsonld.Context) (x schema.EndorseAction) {
	x.Type = "http://schema.org/EndorseAction"
	MapBasicToReactActionTrait(ctx, &x.ReactActionTrait)
	MapCustomToReactActionTrait(ctx, &x.ReactActionTrait)

	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEndorseActionTrait(ctx, &x.EndorseActionTrait)
	MapCustomToEndorseActionTrait(ctx, &x.EndorseActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEndorseActionTrait(ctx jsonld.Context, x *schema.EndorseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEndorseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEndorseActionTrait(ctx jsonld.Context, x *schema.EndorseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEndorseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}