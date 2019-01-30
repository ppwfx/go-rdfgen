package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWebPageElement strings.Replacer
var strconvWebPageElement strconv.NumError

var basicWebPageElementTraitMapping = map[string]func(*schema.WebPageElementTrait, []string){}
var customWebPageElementTraitMapping = map[string]func(*schema.WebPageElementTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WebPageElement", func(ctx jsonld.Context) (interface{}) {
		return NewWebPageElementFromContext(ctx)
	})




	customWebPageElementTraitMapping["http://schema.org/cssSelector"] = func(x *schema.WebPageElementTrait, ctx jsonld.Context, s string){
		var y schema.CssSelectorType
		if strings.HasPrefix(s, "_:") {
			y = NewCssSelectorTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCssSelectorType()
			y.Id = s
		}

		x.CssSelector = &y

		return
	}

	customWebPageElementTraitMapping["http://schema.org/xpath"] = func(x *schema.WebPageElementTrait, ctx jsonld.Context, s string){
		var y schema.XPathType
		if strings.HasPrefix(s, "_:") {
			y = NewXPathTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewXPathType()
			y.Id = s
		}

		x.Xpath = &y

		return
	}
}

func NewWebPageElementFromContext(ctx jsonld.Context) (x schema.WebPageElement) {
	x.Type = "http://schema.org/WebPageElement"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWebPageElementTrait(ctx, &x.WebPageElementTrait)
	MapCustomToWebPageElementTrait(ctx, &x.WebPageElementTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWebPageElementTrait(ctx jsonld.Context, x *schema.WebPageElementTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWebPageElementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWebPageElementTrait(ctx jsonld.Context, x *schema.WebPageElementTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWebPageElementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}