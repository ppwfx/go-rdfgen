package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEnumeration strings.Replacer
var strconvEnumeration strconv.NumError

var basicEnumerationTraitMapping = map[string]func(*schema.EnumerationTrait, []string){}
var customEnumerationTraitMapping = map[string]func(*schema.EnumerationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Enumeration", func(ctx jsonld.Context) (interface{}) {
		return NewEnumerationFromContext(ctx)
	})



	customEnumerationTraitMapping["http://schema.org/supersededBy"] = func(x *schema.EnumerationTrait, ctx jsonld.Context, s string){
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

func NewEnumerationFromContext(ctx jsonld.Context) (x schema.Enumeration) {
	x.Type = "http://schema.org/Enumeration"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEnumerationTrait(ctx jsonld.Context, x *schema.EnumerationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEnumerationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEnumerationTrait(ctx jsonld.Context, x *schema.EnumerationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEnumerationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}