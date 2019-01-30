package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLanguage strings.Replacer
var strconvLanguage strconv.NumError

var basicLanguageTraitMapping = map[string]func(*schema.LanguageTrait, []string){}
var customLanguageTraitMapping = map[string]func(*schema.LanguageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Language", func(ctx jsonld.Context) (interface{}) {
		return NewLanguageFromContext(ctx)
	})

}

func NewLanguageFromContext(ctx jsonld.Context) (x schema.Language) {
	x.Type = "http://schema.org/Language"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLanguageTrait(ctx, &x.LanguageTrait)
	MapCustomToLanguageTrait(ctx, &x.LanguageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLanguageTrait(ctx jsonld.Context, x *schema.LanguageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLanguageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLanguageTrait(ctx jsonld.Context, x *schema.LanguageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLanguageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}