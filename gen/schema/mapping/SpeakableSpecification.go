package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSpeakableSpecification strings.Replacer
var strconvSpeakableSpecification strconv.NumError

var basicSpeakableSpecificationTraitMapping = map[string]func(*schema.SpeakableSpecificationTrait, []string){}
var customSpeakableSpecificationTraitMapping = map[string]func(*schema.SpeakableSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SpeakableSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewSpeakableSpecificationFromContext(ctx)
	})




	customSpeakableSpecificationTraitMapping["http://schema.org/cssSelector"] = func(x *schema.SpeakableSpecificationTrait, ctx jsonld.Context, s string){
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

	customSpeakableSpecificationTraitMapping["http://schema.org/xpath"] = func(x *schema.SpeakableSpecificationTrait, ctx jsonld.Context, s string){
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

func NewSpeakableSpecificationFromContext(ctx jsonld.Context) (x schema.SpeakableSpecification) {
	x.Type = "http://schema.org/SpeakableSpecification"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSpeakableSpecificationTrait(ctx, &x.SpeakableSpecificationTrait)
	MapCustomToSpeakableSpecificationTrait(ctx, &x.SpeakableSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSpeakableSpecificationTrait(ctx jsonld.Context, x *schema.SpeakableSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSpeakableSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSpeakableSpecificationTrait(ctx jsonld.Context, x *schema.SpeakableSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSpeakableSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}