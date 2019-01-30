package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsClass strings.Replacer
var strconvClass strconv.NumError

var basicClassTraitMapping = map[string]func(*schema.ClassTrait, []string){}
var customClassTraitMapping = map[string]func(*schema.ClassTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Class", func(ctx jsonld.Context) (interface{}) {
		return NewClassFromContext(ctx)
	})



	customClassTraitMapping["http://schema.org/supersededBy"] = func(x *schema.ClassTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SupersededBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SupersededBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SupersededBy = s
		}
	}
}

func NewClassFromContext(ctx jsonld.Context) (x schema.Class) {
	x.Type = "http://schema.org/Class"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToClassTrait(ctx, &x.ClassTrait)
	MapCustomToClassTrait(ctx, &x.ClassTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToClassTrait(ctx jsonld.Context, x *schema.ClassTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicClassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToClassTrait(ctx jsonld.Context, x *schema.ClassTrait) () {
	for k, v := range ctx.Current {
		f, ok := customClassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}