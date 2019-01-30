package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQuantitativeValueDistribution strings.Replacer
var strconvQuantitativeValueDistribution strconv.NumError

var basicQuantitativeValueDistributionTraitMapping = map[string]func(*schema.QuantitativeValueDistributionTrait, []string){}
var customQuantitativeValueDistributionTraitMapping = map[string]func(*schema.QuantitativeValueDistributionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/QuantitativeValueDistribution", func(ctx jsonld.Context) (interface{}) {
		return NewQuantitativeValueDistributionFromContext(ctx)
	})



	basicQuantitativeValueDistributionTraitMapping["http://schema.org/median"] = func(x *schema.QuantitativeValueDistributionTrait, s []string) {
		var err error
		x.Median, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicQuantitativeValueDistributionTraitMapping["http://schema.org/percentile10"] = func(x *schema.QuantitativeValueDistributionTrait, s []string) {
		var err error
		x.Percentile10, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicQuantitativeValueDistributionTraitMapping["http://schema.org/percentile25"] = func(x *schema.QuantitativeValueDistributionTrait, s []string) {
		var err error
		x.Percentile25, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicQuantitativeValueDistributionTraitMapping["http://schema.org/percentile75"] = func(x *schema.QuantitativeValueDistributionTrait, s []string) {
		var err error
		x.Percentile75, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicQuantitativeValueDistributionTraitMapping["http://schema.org/percentile90"] = func(x *schema.QuantitativeValueDistributionTrait, s []string) {
		var err error
		x.Percentile90, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	customQuantitativeValueDistributionTraitMapping["http://schema.org/duration"] = func(x *schema.QuantitativeValueDistributionTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.Duration = &y

		return
	}
}

func NewQuantitativeValueDistributionFromContext(ctx jsonld.Context) (x schema.QuantitativeValueDistribution) {
	x.Type = "http://schema.org/QuantitativeValueDistribution"
	MapBasicToQuantitativeValueTrait(ctx, &x.QuantitativeValueTrait)
	MapCustomToQuantitativeValueTrait(ctx, &x.QuantitativeValueTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQuantitativeValueDistributionTrait(ctx, &x.QuantitativeValueDistributionTrait)
	MapCustomToQuantitativeValueDistributionTrait(ctx, &x.QuantitativeValueDistributionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQuantitativeValueDistributionTrait(ctx jsonld.Context, x *schema.QuantitativeValueDistributionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQuantitativeValueDistributionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQuantitativeValueDistributionTrait(ctx jsonld.Context, x *schema.QuantitativeValueDistributionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQuantitativeValueDistributionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}