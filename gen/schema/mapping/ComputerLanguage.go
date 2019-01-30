package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComputerLanguage strings.Replacer
var strconvComputerLanguage strconv.NumError

var basicComputerLanguageTraitMapping = map[string]func(*schema.ComputerLanguageTrait, []string){}
var customComputerLanguageTraitMapping = map[string]func(*schema.ComputerLanguageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComputerLanguage", func(ctx jsonld.Context) (interface{}) {
		return NewComputerLanguageFromContext(ctx)
	})

}

func NewComputerLanguageFromContext(ctx jsonld.Context) (x schema.ComputerLanguage) {
	x.Type = "http://schema.org/ComputerLanguage"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToComputerLanguageTrait(ctx, &x.ComputerLanguageTrait)
	MapCustomToComputerLanguageTrait(ctx, &x.ComputerLanguageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComputerLanguageTrait(ctx jsonld.Context, x *schema.ComputerLanguageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComputerLanguageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComputerLanguageTrait(ctx jsonld.Context, x *schema.ComputerLanguageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComputerLanguageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}