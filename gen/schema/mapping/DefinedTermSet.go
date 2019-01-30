package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDefinedTermSet strings.Replacer
var strconvDefinedTermSet strconv.NumError

var basicDefinedTermSetTraitMapping = map[string]func(*schema.DefinedTermSetTrait, []string){}
var customDefinedTermSetTraitMapping = map[string]func(*schema.DefinedTermSetTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DefinedTermSet", func(ctx jsonld.Context) (interface{}) {
		return NewDefinedTermSetFromContext(ctx)
	})



	customDefinedTermSetTraitMapping["http://schema.org/hasDefinedTerm"] = func(x *schema.DefinedTermSetTrait, ctx jsonld.Context, s string){
		var y schema.DefinedTerm
		if strings.HasPrefix(s, "_:") {
			y = NewDefinedTermFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDefinedTerm()
			y.Id = s
		}

		x.HasDefinedTerm = &y

		return
	}
}

func NewDefinedTermSetFromContext(ctx jsonld.Context) (x schema.DefinedTermSet) {
	x.Type = "http://schema.org/DefinedTermSet"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDefinedTermSetTrait(ctx, &x.DefinedTermSetTrait)
	MapCustomToDefinedTermSetTrait(ctx, &x.DefinedTermSetTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDefinedTermSetTrait(ctx jsonld.Context, x *schema.DefinedTermSetTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDefinedTermSetTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDefinedTermSetTrait(ctx jsonld.Context, x *schema.DefinedTermSetTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDefinedTermSetTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}