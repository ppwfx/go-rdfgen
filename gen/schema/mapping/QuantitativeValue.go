package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQuantitativeValue strings.Replacer
var strconvQuantitativeValue strconv.NumError

var basicQuantitativeValueTraitMapping = map[string]func(*schema.QuantitativeValueTrait, []string){}
var customQuantitativeValueTraitMapping = map[string]func(*schema.QuantitativeValueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/QuantitativeValue", func(ctx jsonld.Context) (interface{}) {
		return NewQuantitativeValueFromContext(ctx)
	})


	basicQuantitativeValueTraitMapping["http://schema.org/maxValue"] = func(x *schema.QuantitativeValueTrait, s []string) {
		var err error
		x.MaxValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicQuantitativeValueTraitMapping["http://schema.org/minValue"] = func(x *schema.QuantitativeValueTrait, s []string) {
		var err error
		x.MinValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}




	basicQuantitativeValueTraitMapping["http://schema.org/unitText"] = func(x *schema.QuantitativeValueTrait, s []string) {
		x.UnitText = s[0]
	}




	customQuantitativeValueTraitMapping["http://schema.org/value"] = func(x *schema.QuantitativeValueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Value, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Value = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Value = s
		}
	}

	customQuantitativeValueTraitMapping["http://schema.org/unitCode"] = func(x *schema.QuantitativeValueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.UnitCode, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.UnitCode = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.UnitCode = s
		}
	}

	customQuantitativeValueTraitMapping["http://schema.org/valueReference"] = func(x *schema.QuantitativeValueTrait, ctx jsonld.Context, s string){
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

	customQuantitativeValueTraitMapping["http://schema.org/additionalProperty"] = func(x *schema.QuantitativeValueTrait, ctx jsonld.Context, s string){
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
}

func NewQuantitativeValueFromContext(ctx jsonld.Context) (x schema.QuantitativeValue) {
	x.Type = "http://schema.org/QuantitativeValue"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQuantitativeValueTrait(ctx, &x.QuantitativeValueTrait)
	MapCustomToQuantitativeValueTrait(ctx, &x.QuantitativeValueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQuantitativeValueTrait(ctx jsonld.Context, x *schema.QuantitativeValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQuantitativeValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQuantitativeValueTrait(ctx jsonld.Context, x *schema.QuantitativeValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQuantitativeValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}