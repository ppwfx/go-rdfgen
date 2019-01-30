package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDefinedTerm strings.Replacer
var strconvDefinedTerm strconv.NumError

var basicDefinedTermTraitMapping = map[string]func(*schema.DefinedTermTrait, []string){}
var customDefinedTermTraitMapping = map[string]func(*schema.DefinedTermTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DefinedTerm", func(ctx jsonld.Context) (interface{}) {
		return NewDefinedTermFromContext(ctx)
	})



	basicDefinedTermTraitMapping["http://schema.org/termCode"] = func(x *schema.DefinedTermTrait, s []string) {
		x.TermCode = s[0]
	}


	customDefinedTermTraitMapping["http://schema.org/inDefinedTermSet"] = func(x *schema.DefinedTermTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.InDefinedTermSet, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.InDefinedTermSet = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.InDefinedTermSet = s
		}
	}
}

func NewDefinedTermFromContext(ctx jsonld.Context) (x schema.DefinedTerm) {
	x.Type = "http://schema.org/DefinedTerm"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDefinedTermTrait(ctx, &x.DefinedTermTrait)
	MapCustomToDefinedTermTrait(ctx, &x.DefinedTermTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDefinedTermTrait(ctx jsonld.Context, x *schema.DefinedTermTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDefinedTermTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDefinedTermTrait(ctx jsonld.Context, x *schema.DefinedTermTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDefinedTermTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}