package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSchedule strings.Replacer
var strconvSchedule strconv.NumError

var basicScheduleTraitMapping = map[string]func(*schema.ScheduleTrait, []string){}
var customScheduleTraitMapping = map[string]func(*schema.ScheduleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Schedule", func(ctx jsonld.Context) (interface{}) {
		return NewScheduleFromContext(ctx)
	})









	customScheduleTraitMapping["http://schema.org/byDay"] = func(x *schema.ScheduleTrait, ctx jsonld.Context, s string){
		var y schema.DayOfWeek
		if strings.HasPrefix(s, "_:") {
			y = NewDayOfWeekFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDayOfWeek()
			y.Id = s
		}

		x.ByDay = &y

		return
	}

	customScheduleTraitMapping["http://schema.org/byMonth"] = func(x *schema.ScheduleTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.ByMonth = &y

		return
	}

	customScheduleTraitMapping["http://schema.org/byMonthDay"] = func(x *schema.ScheduleTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.ByMonthDay = &y

		return
	}

	customScheduleTraitMapping["http://schema.org/eventSchedule"] = func(x *schema.ScheduleTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.EventSchedule = &y

		return
	}

	customScheduleTraitMapping["http://schema.org/exceptDate"] = func(x *schema.ScheduleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ExceptDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ExceptDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ExceptDate = s
		}
	}

	customScheduleTraitMapping["http://schema.org/repeatCount"] = func(x *schema.ScheduleTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.RepeatCount = &y

		return
	}

	customScheduleTraitMapping["http://schema.org/repeatFrequency"] = func(x *schema.ScheduleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RepeatFrequency, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RepeatFrequency = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RepeatFrequency = s
		}
	}
}

func NewScheduleFromContext(ctx jsonld.Context) (x schema.Schedule) {
	x.Type = "http://schema.org/Schedule"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToScheduleTrait(ctx, &x.ScheduleTrait)
	MapCustomToScheduleTrait(ctx, &x.ScheduleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToScheduleTrait(ctx jsonld.Context, x *schema.ScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToScheduleTrait(ctx jsonld.Context, x *schema.ScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}