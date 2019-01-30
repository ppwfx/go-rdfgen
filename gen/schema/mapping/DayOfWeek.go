package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDayOfWeek strings.Replacer
var strconvDayOfWeek strconv.NumError

var basicDayOfWeekTraitMapping = map[string]func(*schema.DayOfWeekTrait, []string){}
var customDayOfWeekTraitMapping = map[string]func(*schema.DayOfWeekTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DayOfWeek", func(ctx jsonld.Context) (interface{}) {
		return NewDayOfWeekFromContext(ctx)
	})

}

func NewDayOfWeekFromContext(ctx jsonld.Context) (x schema.DayOfWeek) {
	x.Type = "http://schema.org/DayOfWeek"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDayOfWeekTrait(ctx, &x.DayOfWeekTrait)
	MapCustomToDayOfWeekTrait(ctx, &x.DayOfWeekTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDayOfWeekTrait(ctx jsonld.Context, x *schema.DayOfWeekTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDayOfWeekTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDayOfWeekTrait(ctx jsonld.Context, x *schema.DayOfWeekTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDayOfWeekTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}