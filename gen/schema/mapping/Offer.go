package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOffer strings.Replacer
var strconvOffer strconv.NumError

var basicOfferTraitMapping = map[string]func(*schema.OfferTrait, []string){}
var customOfferTraitMapping = map[string]func(*schema.OfferTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Offer", func(ctx jsonld.Context) (interface{}) {
		return NewOfferFromContext(ctx)
	})






















	basicOfferTraitMapping["http://schema.org/gtin12"] = func(x *schema.OfferTrait, s []string) {
		x.Gtin12 = s[0]
	}


	basicOfferTraitMapping["http://schema.org/gtin13"] = func(x *schema.OfferTrait, s []string) {
		x.Gtin13 = s[0]
	}


	basicOfferTraitMapping["http://schema.org/gtin14"] = func(x *schema.OfferTrait, s []string) {
		x.Gtin14 = s[0]
	}


	basicOfferTraitMapping["http://schema.org/gtin8"] = func(x *schema.OfferTrait, s []string) {
		x.Gtin8 = s[0]
	}







	basicOfferTraitMapping["http://schema.org/mpn"] = func(x *schema.OfferTrait, s []string) {
		x.Mpn = s[0]
	}




	basicOfferTraitMapping["http://schema.org/serialNumber"] = func(x *schema.OfferTrait, s []string) {
		x.SerialNumber = s[0]
	}


	basicOfferTraitMapping["http://schema.org/sku"] = func(x *schema.OfferTrait, s []string) {
		x.Sku = s[0]
	}





	basicOfferTraitMapping["http://schema.org/priceCurrency"] = func(x *schema.OfferTrait, s []string) {
		x.PriceCurrency = s[0]
	}





	customOfferTraitMapping["http://schema.org/validFrom"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
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

	customOfferTraitMapping["http://schema.org/validThrough"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
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

	customOfferTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
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

	customOfferTraitMapping["http://schema.org/review"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
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

	customOfferTraitMapping["http://schema.org/reviews"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
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

	customOfferTraitMapping["http://schema.org/acceptedPaymentMethod"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AcceptedPaymentMethod, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AcceptedPaymentMethod = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AcceptedPaymentMethod = s
		}
	}

	customOfferTraitMapping["http://schema.org/advanceBookingRequirement"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.AdvanceBookingRequirement = &y

		return
	}

	customOfferTraitMapping["http://schema.org/areaServed"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
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

	customOfferTraitMapping["http://schema.org/availability"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.ItemAvailability
		if strings.HasPrefix(s, "_:") {
			y = NewItemAvailabilityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewItemAvailability()
			y.Id = s
		}

		x.Availability = &y

		return
	}

	customOfferTraitMapping["http://schema.org/availabilityEnds"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.AvailabilityEnds = &y

		return
	}

	customOfferTraitMapping["http://schema.org/availabilityStarts"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.AvailabilityStarts = &y

		return
	}

	customOfferTraitMapping["http://schema.org/availableAtOrFrom"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.AvailableAtOrFrom = &y

		return
	}

	customOfferTraitMapping["http://schema.org/availableDeliveryMethod"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.DeliveryMethod
		if strings.HasPrefix(s, "_:") {
			y = NewDeliveryMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDeliveryMethod()
			y.Id = s
		}

		x.AvailableDeliveryMethod = &y

		return
	}

	customOfferTraitMapping["http://schema.org/businessFunction"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.BusinessFunction
		if strings.HasPrefix(s, "_:") {
			y = NewBusinessFunctionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBusinessFunction()
			y.Id = s
		}

		x.BusinessFunction = &y

		return
	}

	customOfferTraitMapping["http://schema.org/deliveryLeadTime"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.DeliveryLeadTime = &y

		return
	}

	customOfferTraitMapping["http://schema.org/eligibleCustomerType"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.BusinessEntityType
		if strings.HasPrefix(s, "_:") {
			y = NewBusinessEntityTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBusinessEntityType()
			y.Id = s
		}

		x.EligibleCustomerType = &y

		return
	}

	customOfferTraitMapping["http://schema.org/eligibleDuration"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.EligibleDuration = &y

		return
	}

	customOfferTraitMapping["http://schema.org/eligibleQuantity"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.EligibleQuantity = &y

		return
	}

	customOfferTraitMapping["http://schema.org/eligibleRegion"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EligibleRegion, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EligibleRegion = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EligibleRegion = s
		}
	}

	customOfferTraitMapping["http://schema.org/eligibleTransactionVolume"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.PriceSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewPriceSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPriceSpecification()
			y.Id = s
		}

		x.EligibleTransactionVolume = &y

		return
	}

	customOfferTraitMapping["http://schema.org/includesObject"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.TypeAndQuantityNode
		if strings.HasPrefix(s, "_:") {
			y = NewTypeAndQuantityNodeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTypeAndQuantityNode()
			y.Id = s
		}

		x.IncludesObject = &y

		return
	}

	customOfferTraitMapping["http://schema.org/ineligibleRegion"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IneligibleRegion, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IneligibleRegion = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IneligibleRegion = s
		}
	}

	customOfferTraitMapping["http://schema.org/inventoryLevel"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.InventoryLevel = &y

		return
	}

	customOfferTraitMapping["http://schema.org/itemCondition"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.OfferItemCondition
		if strings.HasPrefix(s, "_:") {
			y = NewOfferItemConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOfferItemCondition()
			y.Id = s
		}

		x.ItemCondition = &y

		return
	}

	customOfferTraitMapping["http://schema.org/itemOffered"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ItemOffered, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ItemOffered = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ItemOffered = s
		}
	}

	customOfferTraitMapping["http://schema.org/priceSpecification"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.PriceSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewPriceSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPriceSpecification()
			y.Id = s
		}

		x.PriceSpecification = &y

		return
	}

	customOfferTraitMapping["http://schema.org/seller"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Seller, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Seller = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Seller = s
		}
	}

	customOfferTraitMapping["http://schema.org/warranty"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.WarrantyPromise
		if strings.HasPrefix(s, "_:") {
			y = NewWarrantyPromiseFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewWarrantyPromise()
			y.Id = s
		}

		x.Warranty = &y

		return
	}

	customOfferTraitMapping["http://schema.org/category"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
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

	customOfferTraitMapping["http://schema.org/price"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Price, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Price = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Price = s
		}
	}

	customOfferTraitMapping["http://schema.org/addOn"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.AddOn = &y

		return
	}

	customOfferTraitMapping["http://schema.org/offeredBy"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.OfferedBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.OfferedBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.OfferedBy = s
		}
	}

	customOfferTraitMapping["http://schema.org/priceValidUntil"] = func(x *schema.OfferTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.PriceValidUntil = &y

		return
	}
}

func NewOfferFromContext(ctx jsonld.Context) (x schema.Offer) {
	x.Type = "http://schema.org/Offer"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOfferTrait(ctx, &x.OfferTrait)
	MapCustomToOfferTrait(ctx, &x.OfferTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOfferTrait(ctx jsonld.Context, x *schema.OfferTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOfferTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOfferTrait(ctx jsonld.Context, x *schema.OfferTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOfferTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}