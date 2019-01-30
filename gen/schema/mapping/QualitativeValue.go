package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQualitativeValue strings.Replacer
var strconvQualitativeValue strconv.NumError

var basicQualitativeValueTraitMapping = map[string]func(*schema.QualitativeValueTrait, []string){}
var customQualitativeValueTraitMapping = map[string]func(*schema.QualitativeValueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/QualitativeValue", func(ctx jsonld.Context) (interface{}) {
		return NewQualitativeValueFromContext(ctx)
	})










	customQualitativeValueTraitMapping["http://schema.org/valueReference"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ValueReference, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ValueReference = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ValueReference = s
		}
	}

	customQualitativeValueTraitMapping["http://schema.org/additionalProperty"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		var y schema.PropertyValue
		if strings.HasPrefix(s, "_:") {
			y = NewPropertyValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPropertyValue()
			y.Id = s
		}

		x.AdditionalProperty = &y

		return
	}

	customQualitativeValueTraitMapping["http://schema.org/equal"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		var y schema.QualitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQualitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQualitativeValue()
			y.Id = s
		}

		x.Equal = &y

		return
	}

	customQualitativeValueTraitMapping["http://schema.org/greater"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		var y schema.QualitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQualitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQualitativeValue()
			y.Id = s
		}

		x.Greater = &y

		return
	}

	customQualitativeValueTraitMapping["http://schema.org/greaterOrEqual"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		var y schema.QualitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQualitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQualitativeValue()
			y.Id = s
		}

		x.GreaterOrEqual = &y

		return
	}

	customQualitativeValueTraitMapping["http://schema.org/lesser"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		var y schema.QualitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQualitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQualitativeValue()
			y.Id = s
		}

		x.Lesser = &y

		return
	}

	customQualitativeValueTraitMapping["http://schema.org/lesserOrEqual"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		var y schema.QualitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQualitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQualitativeValue()
			y.Id = s
		}

		x.LesserOrEqual = &y

		return
	}

	customQualitativeValueTraitMapping["http://schema.org/nonEqual"] = func(x *schema.QualitativeValueTrait, ctx jsonld.Context, s string){
		var y schema.QualitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQualitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQualitativeValue()
			y.Id = s
		}

		x.NonEqual = &y

		return
	}
}

func NewQualitativeValueFromContext(ctx jsonld.Context) (x schema.QualitativeValue) {
	x.Type = "http://schema.org/QualitativeValue"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)
	MapCustomToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQualitativeValueTrait(ctx jsonld.Context, x *schema.QualitativeValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQualitativeValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQualitativeValueTrait(ctx jsonld.Context, x *schema.QualitativeValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQualitativeValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}