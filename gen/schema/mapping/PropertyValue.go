package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPropertyValue strings.Replacer
var strconvPropertyValue strconv.NumError

var basicPropertyValueTraitMapping = map[string]func(*schema.PropertyValueTrait, []string){}
var customPropertyValueTraitMapping = map[string]func(*schema.PropertyValueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PropertyValue", func(ctx jsonld.Context) (interface{}) {
		return NewPropertyValueFromContext(ctx)
	})


	basicPropertyValueTraitMapping["http://schema.org/maxValue"] = func(x *schema.PropertyValueTrait, s []string) {
		var err error
		x.MaxValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicPropertyValueTraitMapping["http://schema.org/minValue"] = func(x *schema.PropertyValueTrait, s []string) {
		var err error
		x.MinValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}






	basicPropertyValueTraitMapping["http://schema.org/unitText"] = func(x *schema.PropertyValueTrait, s []string) {
		x.UnitText = s[0]
	}



	customPropertyValueTraitMapping["http://schema.org/value"] = func(x *schema.PropertyValueTrait, ctx jsonld.Context, s string){
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

	customPropertyValueTraitMapping["http://schema.org/measurementTechnique"] = func(x *schema.PropertyValueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MeasurementTechnique, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MeasurementTechnique = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MeasurementTechnique = s
		}
	}

	customPropertyValueTraitMapping["http://schema.org/propertyID"] = func(x *schema.PropertyValueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PropertyID, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PropertyID = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PropertyID = s
		}
	}

	customPropertyValueTraitMapping["http://schema.org/unitCode"] = func(x *schema.PropertyValueTrait, ctx jsonld.Context, s string){
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

	customPropertyValueTraitMapping["http://schema.org/valueReference"] = func(x *schema.PropertyValueTrait, ctx jsonld.Context, s string){
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
}

func NewPropertyValueFromContext(ctx jsonld.Context) (x schema.PropertyValue) {
	x.Type = "http://schema.org/PropertyValue"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPropertyValueTrait(ctx, &x.PropertyValueTrait)
	MapCustomToPropertyValueTrait(ctx, &x.PropertyValueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPropertyValueTrait(ctx jsonld.Context, x *schema.PropertyValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPropertyValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPropertyValueTrait(ctx jsonld.Context, x *schema.PropertyValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPropertyValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}