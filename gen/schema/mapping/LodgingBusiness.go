package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLodgingBusiness strings.Replacer
var strconvLodgingBusiness strconv.NumError

var basicLodgingBusinessTraitMapping = map[string]func(*schema.LodgingBusinessTrait, []string){}
var customLodgingBusinessTraitMapping = map[string]func(*schema.LodgingBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LodgingBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewLodgingBusinessFromContext(ctx)
	})









	customLodgingBusinessTraitMapping["http://schema.org/audience"] = func(x *schema.LodgingBusinessTrait, ctx jsonld.Context, s string){
		var y schema.Audience
		if strings.HasPrefix(s, "_:") {
			y = NewAudienceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAudience()
			y.Id = s
		}

		x.Audience = &y

		return
	}

	customLodgingBusinessTraitMapping["http://schema.org/availableLanguage"] = func(x *schema.LodgingBusinessTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AvailableLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AvailableLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AvailableLanguage = s
		}
	}

	customLodgingBusinessTraitMapping["http://schema.org/amenityFeature"] = func(x *schema.LodgingBusinessTrait, ctx jsonld.Context, s string){
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

	customLodgingBusinessTraitMapping["http://schema.org/checkinTime"] = func(x *schema.LodgingBusinessTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.CheckinTime = &y

		return
	}

	customLodgingBusinessTraitMapping["http://schema.org/checkoutTime"] = func(x *schema.LodgingBusinessTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.CheckoutTime = &y

		return
	}

	customLodgingBusinessTraitMapping["http://schema.org/petsAllowed"] = func(x *schema.LodgingBusinessTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PetsAllowed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PetsAllowed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PetsAllowed = s
		}
	}

	customLodgingBusinessTraitMapping["http://schema.org/starRating"] = func(x *schema.LodgingBusinessTrait, ctx jsonld.Context, s string){
		var y schema.Rating
		if strings.HasPrefix(s, "_:") {
			y = NewRatingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRating()
			y.Id = s
		}

		x.StarRating = &y

		return
	}
}

func NewLodgingBusinessFromContext(ctx jsonld.Context) (x schema.LodgingBusiness) {
	x.Type = "http://schema.org/LodgingBusiness"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)
	MapCustomToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLodgingBusinessTrait(ctx jsonld.Context, x *schema.LodgingBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLodgingBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLodgingBusinessTrait(ctx jsonld.Context, x *schema.LodgingBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLodgingBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}