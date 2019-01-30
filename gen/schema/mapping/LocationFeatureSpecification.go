package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLocationFeatureSpecification strings.Replacer
var strconvLocationFeatureSpecification strconv.NumError

var basicLocationFeatureSpecificationTraitMapping = map[string]func(*schema.LocationFeatureSpecificationTrait, []string){}
var customLocationFeatureSpecificationTraitMapping = map[string]func(*schema.LocationFeatureSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LocationFeatureSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewLocationFeatureSpecificationFromContext(ctx)
	})





	customLocationFeatureSpecificationTraitMapping["http://schema.org/validFrom"] = func(x *schema.LocationFeatureSpecificationTrait, ctx jsonld.Context, s string){
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

	customLocationFeatureSpecificationTraitMapping["http://schema.org/validThrough"] = func(x *schema.LocationFeatureSpecificationTrait, ctx jsonld.Context, s string){
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

	customLocationFeatureSpecificationTraitMapping["http://schema.org/hoursAvailable"] = func(x *schema.LocationFeatureSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.OpeningHoursSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewOpeningHoursSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOpeningHoursSpecification()
			y.Id = s
		}

		x.HoursAvailable = &y

		return
	}
}

func NewLocationFeatureSpecificationFromContext(ctx jsonld.Context) (x schema.LocationFeatureSpecification) {
	x.Type = "http://schema.org/LocationFeatureSpecification"
	MapBasicToPropertyValueTrait(ctx, &x.PropertyValueTrait)
	MapCustomToPropertyValueTrait(ctx, &x.PropertyValueTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLocationFeatureSpecificationTrait(ctx, &x.LocationFeatureSpecificationTrait)
	MapCustomToLocationFeatureSpecificationTrait(ctx, &x.LocationFeatureSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLocationFeatureSpecificationTrait(ctx jsonld.Context, x *schema.LocationFeatureSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLocationFeatureSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLocationFeatureSpecificationTrait(ctx jsonld.Context, x *schema.LocationFeatureSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLocationFeatureSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}