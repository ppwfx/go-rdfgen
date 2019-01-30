package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPlace strings.Replacer
var strconvPlace strconv.NumError

var basicPlaceTraitMapping = map[string]func(*schema.PlaceTrait, []string){}
var customPlaceTraitMapping = map[string]func(*schema.PlaceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Place", func(ctx jsonld.Context) (interface{}) {
		return NewPlaceFromContext(ctx)
	})



	basicPlaceTraitMapping["http://schema.org/isAccessibleForFree"] = func(x *schema.PlaceTrait, s []string) {
		var err error
		x.IsAccessibleForFree, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}









	basicPlaceTraitMapping["http://schema.org/faxNumber"] = func(x *schema.PlaceTrait, s []string) {
		x.FaxNumber = s[0]
	}


	basicPlaceTraitMapping["http://schema.org/telephone"] = func(x *schema.PlaceTrait, s []string) {
		x.Telephone = s[0]
	}



	basicPlaceTraitMapping["http://schema.org/branchCode"] = func(x *schema.PlaceTrait, s []string) {
		x.BranchCode = s[0]
	}


















	basicPlaceTraitMapping["http://schema.org/globalLocationNumber"] = func(x *schema.PlaceTrait, s []string) {
		x.GlobalLocationNumber = s[0]
	}



	basicPlaceTraitMapping["http://schema.org/isicV4"] = func(x *schema.PlaceTrait, s []string) {
		x.IsicV4 = s[0]
	}



	basicPlaceTraitMapping["http://schema.org/map"] = func(x *schema.PlaceTrait, s []string) {
		x.Map = s[0]
	}


	basicPlaceTraitMapping["http://schema.org/maps"] = func(x *schema.PlaceTrait, s []string) {
		x.Maps = s[0]
	}




	basicPlaceTraitMapping["http://schema.org/publicAccess"] = func(x *schema.PlaceTrait, s []string) {
		var err error
		x.PublicAccess, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicPlaceTraitMapping["http://schema.org/smokingAllowed"] = func(x *schema.PlaceTrait, s []string) {
		var err error
		x.SmokingAllowed, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	customPlaceTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.AggregateRating
		if strings.HasPrefix(s, "_:") {
			y = NewAggregateRatingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAggregateRating()
			y.Id = s
		}

		x.AggregateRating = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/review"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.Review = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/reviews"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.Reviews = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/maximumAttendeeCapacity"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.MaximumAttendeeCapacity = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/specialOpeningHoursSpecification"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.OpeningHoursSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewOpeningHoursSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOpeningHoursSpecification()
			y.Id = s
		}

		x.SpecialOpeningHoursSpecification = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/openingHoursSpecification"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.OpeningHoursSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewOpeningHoursSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOpeningHoursSpecification()
			y.Id = s
		}

		x.OpeningHoursSpecification = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/address"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Address, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Address = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Address = s
		}
	}

	customPlaceTraitMapping["http://schema.org/additionalProperty"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.PropertyValue
		if strings.HasPrefix(s, "_:") {
			y = NewPropertyValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPropertyValue()
			y.Id = s
		}

		x.AdditionalProperty = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/amenityFeature"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.LocationFeatureSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewLocationFeatureSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLocationFeatureSpecification()
			y.Id = s
		}

		x.AmenityFeature = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/containedIn"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ContainedIn = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/containedInPlace"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ContainedInPlace = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/containsPlace"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ContainsPlace = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/event"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.Event = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/events"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.Events = &y

		return
	}

	customPlaceTraitMapping["http://schema.org/geo"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Geo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Geo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Geo = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyContains"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyContains, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyContains = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyContains = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyCoveredBy"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyCoveredBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyCoveredBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyCoveredBy = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyCovers"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyCovers, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyCovers = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyCovers = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyCrosses"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyCrosses, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyCrosses = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyCrosses = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyDisjoint"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyDisjoint, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyDisjoint = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyDisjoint = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyEquals"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyEquals, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyEquals = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyEquals = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyIntersects"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyIntersects, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyIntersects = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyIntersects = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyOverlaps"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyOverlaps, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyOverlaps = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyOverlaps = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyTouches"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyTouches, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyTouches = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyTouches = s
		}
	}

	customPlaceTraitMapping["http://schema.org/geospatiallyWithin"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyWithin, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyWithin = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyWithin = s
		}
	}

	customPlaceTraitMapping["http://schema.org/hasMap"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.HasMap, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.HasMap = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.HasMap = s
		}
	}

	customPlaceTraitMapping["http://schema.org/logo"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Logo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Logo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Logo = s
		}
	}

	customPlaceTraitMapping["http://schema.org/photo"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Photo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Photo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Photo = s
		}
	}

	customPlaceTraitMapping["http://schema.org/photos"] = func(x *schema.PlaceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Photos, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Photos = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Photos = s
		}
	}
}

func NewPlaceFromContext(ctx jsonld.Context) (x schema.Place) {
	x.Type = "http://schema.org/Place"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPlaceTrait(ctx jsonld.Context, x *schema.PlaceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPlaceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPlaceTrait(ctx jsonld.Context, x *schema.PlaceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPlaceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}