package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDemand strings.Replacer
var strconvDemand strconv.NumError

var basicDemandTraitMapping = map[string]func(*schema.DemandTrait, []string){}
var customDemandTraitMapping = map[string]func(*schema.DemandTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Demand", func(ctx jsonld.Context) (interface{}) {
		return NewDemandFromContext(ctx)
	})



















	basicDemandTraitMapping["http://schema.org/gtin12"] = func(x *schema.DemandTrait, s []string) {
		x.Gtin12 = s[0]
	}


	basicDemandTraitMapping["http://schema.org/gtin13"] = func(x *schema.DemandTrait, s []string) {
		x.Gtin13 = s[0]
	}


	basicDemandTraitMapping["http://schema.org/gtin14"] = func(x *schema.DemandTrait, s []string) {
		x.Gtin14 = s[0]
	}


	basicDemandTraitMapping["http://schema.org/gtin8"] = func(x *schema.DemandTrait, s []string) {
		x.Gtin8 = s[0]
	}







	basicDemandTraitMapping["http://schema.org/mpn"] = func(x *schema.DemandTrait, s []string) {
		x.Mpn = s[0]
	}




	basicDemandTraitMapping["http://schema.org/serialNumber"] = func(x *schema.DemandTrait, s []string) {
		x.SerialNumber = s[0]
	}


	basicDemandTraitMapping["http://schema.org/sku"] = func(x *schema.DemandTrait, s []string) {
		x.Sku = s[0]
	}



	customDemandTraitMapping["http://schema.org/validFrom"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/validThrough"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/acceptedPaymentMethod"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/advanceBookingRequirement"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/areaServed"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/availability"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/availabilityEnds"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/availabilityStarts"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/availableAtOrFrom"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/availableDeliveryMethod"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/businessFunction"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/deliveryLeadTime"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/eligibleCustomerType"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/eligibleDuration"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/eligibleQuantity"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/eligibleRegion"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/eligibleTransactionVolume"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/includesObject"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/ineligibleRegion"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/inventoryLevel"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/itemCondition"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/itemOffered"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/priceSpecification"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/seller"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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

	customDemandTraitMapping["http://schema.org/warranty"] = func(x *schema.DemandTrait, ctx jsonld.Context, s string){
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
}

func NewDemandFromContext(ctx jsonld.Context) (x schema.Demand) {
	x.Type = "http://schema.org/Demand"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDemandTrait(ctx, &x.DemandTrait)
	MapCustomToDemandTrait(ctx, &x.DemandTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDemandTrait(ctx jsonld.Context, x *schema.DemandTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDemandTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDemandTrait(ctx jsonld.Context, x *schema.DemandTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDemandTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}