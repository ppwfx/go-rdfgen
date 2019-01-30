package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPropertyValueSpecification strings.Replacer
var strconvPropertyValueSpecification strconv.NumError

var basicPropertyValueSpecificationTraitMapping = map[string]func(*schema.PropertyValueSpecificationTrait, []string){}
var customPropertyValueSpecificationTraitMapping = map[string]func(*schema.PropertyValueSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PropertyValueSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewPropertyValueSpecificationFromContext(ctx)
	})


	basicPropertyValueSpecificationTraitMapping["http://schema.org/maxValue"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.MaxValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/minValue"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.MinValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	basicPropertyValueSpecificationTraitMapping["http://schema.org/multipleValues"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.MultipleValues, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/readonlyValue"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.ReadonlyValue, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/stepValue"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.StepValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/valueMaxLength"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.ValueMaxLength, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/valueMinLength"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.ValueMinLength, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/valueName"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		x.ValueName = s[0]
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/valuePattern"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		x.ValuePattern = s[0]
	}


	basicPropertyValueSpecificationTraitMapping["http://schema.org/valueRequired"] = func(x *schema.PropertyValueSpecificationTrait, s []string) {
		var err error
		x.ValueRequired, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	customPropertyValueSpecificationTraitMapping["http://schema.org/defaultValue"] = func(x *schema.PropertyValueSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DefaultValue, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DefaultValue = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DefaultValue = s
		}
	}
}

func NewPropertyValueSpecificationFromContext(ctx jsonld.Context) (x schema.PropertyValueSpecification) {
	x.Type = "http://schema.org/PropertyValueSpecification"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPropertyValueSpecificationTrait(ctx, &x.PropertyValueSpecificationTrait)
	MapCustomToPropertyValueSpecificationTrait(ctx, &x.PropertyValueSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPropertyValueSpecificationTrait(ctx jsonld.Context, x *schema.PropertyValueSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPropertyValueSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPropertyValueSpecificationTrait(ctx jsonld.Context, x *schema.PropertyValueSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPropertyValueSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}