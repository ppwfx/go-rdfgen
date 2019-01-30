package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsService strings.Replacer
var strconvService strconv.NumError

var basicServiceTraitMapping = map[string]func(*schema.ServiceTrait, []string){}
var customServiceTraitMapping = map[string]func(*schema.ServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Service", func(ctx jsonld.Context) (interface{}) {
		return NewServiceFromContext(ctx)
	})




	basicServiceTraitMapping["http://schema.org/award"] = func(x *schema.ServiceTrait, s []string) {
		x.Award = s[0]
	}

















	basicServiceTraitMapping["http://schema.org/providerMobility"] = func(x *schema.ServiceTrait, s []string) {
		x.ProviderMobility = s[0]
	}




	basicServiceTraitMapping["http://schema.org/serviceType"] = func(x *schema.ServiceTrait, s []string) {
		x.ServiceType = s[0]
	}



	customServiceTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
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

	customServiceTraitMapping["http://schema.org/audience"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
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

	customServiceTraitMapping["http://schema.org/offers"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.Offers = &y

		return
	}

	customServiceTraitMapping["http://schema.org/provider"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Provider, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Provider = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Provider = s
		}
	}

	customServiceTraitMapping["http://schema.org/review"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
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

	customServiceTraitMapping["http://schema.org/hoursAvailable"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
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

	customServiceTraitMapping["http://schema.org/areaServed"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AreaServed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AreaServed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AreaServed = s
		}
	}

	customServiceTraitMapping["http://schema.org/category"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Category, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Category = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Category = s
		}
	}

	customServiceTraitMapping["http://schema.org/serviceArea"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ServiceArea, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ServiceArea = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ServiceArea = s
		}
	}

	customServiceTraitMapping["http://schema.org/logo"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
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

	customServiceTraitMapping["http://schema.org/brand"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Brand, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Brand = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Brand = s
		}
	}

	customServiceTraitMapping["http://schema.org/hasOfferCatalog"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		var y schema.OfferCatalog
		if strings.HasPrefix(s, "_:") {
			y = NewOfferCatalogFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOfferCatalog()
			y.Id = s
		}

		x.HasOfferCatalog = &y

		return
	}

	customServiceTraitMapping["http://schema.org/availableChannel"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		var y schema.ServiceChannel
		if strings.HasPrefix(s, "_:") {
			y = NewServiceChannelFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewServiceChannel()
			y.Id = s
		}

		x.AvailableChannel = &y

		return
	}

	customServiceTraitMapping["http://schema.org/broker"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Broker, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Broker = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Broker = s
		}
	}

	customServiceTraitMapping["http://schema.org/isRelatedTo"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsRelatedTo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsRelatedTo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsRelatedTo = s
		}
	}

	customServiceTraitMapping["http://schema.org/isSimilarTo"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsSimilarTo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsSimilarTo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsSimilarTo = s
		}
	}

	customServiceTraitMapping["http://schema.org/produces"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Produces = &y

		return
	}

	customServiceTraitMapping["http://schema.org/serviceAudience"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		var y schema.Audience
		if strings.HasPrefix(s, "_:") {
			y = NewAudienceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAudience()
			y.Id = s
		}

		x.ServiceAudience = &y

		return
	}

	customServiceTraitMapping["http://schema.org/serviceOutput"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.ServiceOutput = &y

		return
	}

	customServiceTraitMapping["http://schema.org/termsOfService"] = func(x *schema.ServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TermsOfService, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TermsOfService = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TermsOfService = s
		}
	}
}

func NewServiceFromContext(ctx jsonld.Context) (x schema.Service) {
	x.Type = "http://schema.org/Service"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToServiceTrait(ctx jsonld.Context, x *schema.ServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToServiceTrait(ctx jsonld.Context, x *schema.ServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}