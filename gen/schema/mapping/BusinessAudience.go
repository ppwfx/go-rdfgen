package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusinessAudience strings.Replacer
var strconvBusinessAudience strconv.NumError

var basicBusinessAudienceTraitMapping = map[string]func(*schema.BusinessAudienceTrait, []string){}
var customBusinessAudienceTraitMapping = map[string]func(*schema.BusinessAudienceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusinessAudience", func(ctx jsonld.Context) (interface{}) {
		return NewBusinessAudienceFromContext(ctx)
	})





	customBusinessAudienceTraitMapping["http://schema.org/numberOfEmployees"] = func(x *schema.BusinessAudienceTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.NumberOfEmployees = &y

		return
	}

	customBusinessAudienceTraitMapping["http://schema.org/yearlyRevenue"] = func(x *schema.BusinessAudienceTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.YearlyRevenue = &y

		return
	}

	customBusinessAudienceTraitMapping["http://schema.org/yearsInOperation"] = func(x *schema.BusinessAudienceTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.YearsInOperation = &y

		return
	}
}

func NewBusinessAudienceFromContext(ctx jsonld.Context) (x schema.BusinessAudience) {
	x.Type = "http://schema.org/BusinessAudience"
	MapBasicToAudienceTrait(ctx, &x.AudienceTrait)
	MapCustomToAudienceTrait(ctx, &x.AudienceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusinessAudienceTrait(ctx, &x.BusinessAudienceTrait)
	MapCustomToBusinessAudienceTrait(ctx, &x.BusinessAudienceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusinessAudienceTrait(ctx jsonld.Context, x *schema.BusinessAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusinessAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusinessAudienceTrait(ctx jsonld.Context, x *schema.BusinessAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusinessAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}