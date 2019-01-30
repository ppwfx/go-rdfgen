package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEngineSpecification strings.Replacer
var strconvEngineSpecification strconv.NumError

var basicEngineSpecificationTraitMapping = map[string]func(*schema.EngineSpecificationTrait, []string){}
var customEngineSpecificationTraitMapping = map[string]func(*schema.EngineSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EngineSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewEngineSpecificationFromContext(ctx)
	})







	customEngineSpecificationTraitMapping["http://schema.org/engineDisplacement"] = func(x *schema.EngineSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.EngineDisplacement = &y

		return
	}

	customEngineSpecificationTraitMapping["http://schema.org/enginePower"] = func(x *schema.EngineSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.EnginePower = &y

		return
	}

	customEngineSpecificationTraitMapping["http://schema.org/engineType"] = func(x *schema.EngineSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EngineType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EngineType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EngineType = s
		}
	}

	customEngineSpecificationTraitMapping["http://schema.org/fuelType"] = func(x *schema.EngineSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FuelType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FuelType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FuelType = s
		}
	}

	customEngineSpecificationTraitMapping["http://schema.org/torque"] = func(x *schema.EngineSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Torque = &y

		return
	}
}

func NewEngineSpecificationFromContext(ctx jsonld.Context) (x schema.EngineSpecification) {
	x.Type = "http://schema.org/EngineSpecification"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEngineSpecificationTrait(ctx, &x.EngineSpecificationTrait)
	MapCustomToEngineSpecificationTrait(ctx, &x.EngineSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEngineSpecificationTrait(ctx jsonld.Context, x *schema.EngineSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEngineSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEngineSpecificationTrait(ctx jsonld.Context, x *schema.EngineSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEngineSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}