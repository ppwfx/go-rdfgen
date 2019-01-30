package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsParcelDelivery strings.Replacer
var strconvParcelDelivery strconv.NumError

var basicParcelDeliveryTraitMapping = map[string]func(*schema.ParcelDeliveryTrait, []string){}
var customParcelDeliveryTraitMapping = map[string]func(*schema.ParcelDeliveryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ParcelDelivery", func(ctx jsonld.Context) (interface{}) {
		return NewParcelDeliveryFromContext(ctx)
	})












	basicParcelDeliveryTraitMapping["http://schema.org/trackingNumber"] = func(x *schema.ParcelDeliveryTrait, s []string) {
		x.TrackingNumber = s[0]
	}


	basicParcelDeliveryTraitMapping["http://schema.org/trackingUrl"] = func(x *schema.ParcelDeliveryTrait, s []string) {
		x.TrackingUrl = s[0]
	}


	customParcelDeliveryTraitMapping["http://schema.org/provider"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
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

	customParcelDeliveryTraitMapping["http://schema.org/carrier"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Carrier = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/deliveryAddress"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.PostalAddress
		if strings.HasPrefix(s, "_:") {
			y = NewPostalAddressFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPostalAddress()
			y.Id = s
		}

		x.DeliveryAddress = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/deliveryStatus"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.DeliveryEvent
		if strings.HasPrefix(s, "_:") {
			y = NewDeliveryEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDeliveryEvent()
			y.Id = s
		}

		x.DeliveryStatus = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/expectedArrivalFrom"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ExpectedArrivalFrom = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/expectedArrivalUntil"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ExpectedArrivalUntil = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/hasDeliveryMethod"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.DeliveryMethod
		if strings.HasPrefix(s, "_:") {
			y = NewDeliveryMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDeliveryMethod()
			y.Id = s
		}

		x.HasDeliveryMethod = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/itemShipped"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.Product
		if strings.HasPrefix(s, "_:") {
			y = NewProductFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProduct()
			y.Id = s
		}

		x.ItemShipped = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/originAddress"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.PostalAddress
		if strings.HasPrefix(s, "_:") {
			y = NewPostalAddressFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPostalAddress()
			y.Id = s
		}

		x.OriginAddress = &y

		return
	}

	customParcelDeliveryTraitMapping["http://schema.org/partOfOrder"] = func(x *schema.ParcelDeliveryTrait, ctx jsonld.Context, s string){
		var y schema.Order
		if strings.HasPrefix(s, "_:") {
			y = NewOrderFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrder()
			y.Id = s
		}

		x.PartOfOrder = &y

		return
	}
}

func NewParcelDeliveryFromContext(ctx jsonld.Context) (x schema.ParcelDelivery) {
	x.Type = "http://schema.org/ParcelDelivery"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToParcelDeliveryTrait(ctx, &x.ParcelDeliveryTrait)
	MapCustomToParcelDeliveryTrait(ctx, &x.ParcelDeliveryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToParcelDeliveryTrait(ctx jsonld.Context, x *schema.ParcelDeliveryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicParcelDeliveryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToParcelDeliveryTrait(ctx jsonld.Context, x *schema.ParcelDeliveryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customParcelDeliveryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}