package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOpeningHoursSpecification strings.Replacer
var strconvOpeningHoursSpecification strconv.NumError

var basicOpeningHoursSpecificationTraitMapping = map[string]func(*schema.OpeningHoursSpecificationTrait, []string){}
var customOpeningHoursSpecificationTraitMapping = map[string]func(*schema.OpeningHoursSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OpeningHoursSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewOpeningHoursSpecificationFromContext(ctx)
	})







	customOpeningHoursSpecificationTraitMapping["http://schema.org/closes"] = func(x *schema.OpeningHoursSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.Time
		if strings.HasPrefix(s, "_:") {
			y = NewTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTime()
			y.Id = s
		}

		x.Closes = &y

		return
	}

	customOpeningHoursSpecificationTraitMapping["http://schema.org/dayOfWeek"] = func(x *schema.OpeningHoursSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DayOfWeek
		if strings.HasPrefix(s, "_:") {
			y = NewDayOfWeekFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDayOfWeek()
			y.Id = s
		}

		x.DayOfWeek = &y

		return
	}

	customOpeningHoursSpecificationTraitMapping["http://schema.org/opens"] = func(x *schema.OpeningHoursSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.Time
		if strings.HasPrefix(s, "_:") {
			y = NewTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTime()
			y.Id = s
		}

		x.Opens = &y

		return
	}

	customOpeningHoursSpecificationTraitMapping["http://schema.org/validFrom"] = func(x *schema.OpeningHoursSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ValidFrom = &y

		return
	}

	customOpeningHoursSpecificationTraitMapping["http://schema.org/validThrough"] = func(x *schema.OpeningHoursSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ValidThrough = &y

		return
	}
}

func NewOpeningHoursSpecificationFromContext(ctx jsonld.Context) (x schema.OpeningHoursSpecification) {
	x.Type = "http://schema.org/OpeningHoursSpecification"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOpeningHoursSpecificationTrait(ctx, &x.OpeningHoursSpecificationTrait)
	MapCustomToOpeningHoursSpecificationTrait(ctx, &x.OpeningHoursSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOpeningHoursSpecificationTrait(ctx jsonld.Context, x *schema.OpeningHoursSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOpeningHoursSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOpeningHoursSpecificationTrait(ctx jsonld.Context, x *schema.OpeningHoursSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOpeningHoursSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}