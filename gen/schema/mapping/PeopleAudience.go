package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPeopleAudience strings.Replacer
var strconvPeopleAudience strconv.NumError

var basicPeopleAudienceTraitMapping = map[string]func(*schema.PeopleAudienceTrait, []string){}
var customPeopleAudienceTraitMapping = map[string]func(*schema.PeopleAudienceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PeopleAudience", func(ctx jsonld.Context) (interface{}) {
		return NewPeopleAudienceFromContext(ctx)
	})



	basicPeopleAudienceTraitMapping["http://schema.org/requiredGender"] = func(x *schema.PeopleAudienceTrait, s []string) {
		x.RequiredGender = s[0]
	}


	basicPeopleAudienceTraitMapping["http://schema.org/suggestedGender"] = func(x *schema.PeopleAudienceTrait, s []string) {
		x.SuggestedGender = s[0]
	}




	basicPeopleAudienceTraitMapping["http://schema.org/suggestedMaxAge"] = func(x *schema.PeopleAudienceTrait, s []string) {
		var err error
		x.SuggestedMaxAge, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicPeopleAudienceTraitMapping["http://schema.org/suggestedMinAge"] = func(x *schema.PeopleAudienceTrait, s []string) {
		var err error
		x.SuggestedMinAge, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	customPeopleAudienceTraitMapping["http://schema.org/healthCondition"] = func(x *schema.PeopleAudienceTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCondition
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCondition()
			y.Id = s
		}

		x.HealthCondition = &y

		return
	}

	customPeopleAudienceTraitMapping["http://schema.org/requiredMaxAge"] = func(x *schema.PeopleAudienceTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.RequiredMaxAge = &y

		return
	}

	customPeopleAudienceTraitMapping["http://schema.org/requiredMinAge"] = func(x *schema.PeopleAudienceTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.RequiredMinAge = &y

		return
	}
}

func NewPeopleAudienceFromContext(ctx jsonld.Context) (x schema.PeopleAudience) {
	x.Type = "http://schema.org/PeopleAudience"
	MapBasicToAudienceTrait(ctx, &x.AudienceTrait)
	MapCustomToAudienceTrait(ctx, &x.AudienceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)
	MapCustomToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPeopleAudienceTrait(ctx jsonld.Context, x *schema.PeopleAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPeopleAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPeopleAudienceTrait(ctx jsonld.Context, x *schema.PeopleAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPeopleAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}