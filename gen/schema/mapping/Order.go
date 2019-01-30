package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOrder strings.Replacer
var strconvOrder strconv.NumError

var basicOrderTraitMapping = map[string]func(*schema.OrderTrait, []string){}
var customOrderTraitMapping = map[string]func(*schema.OrderTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Order", func(ctx jsonld.Context) (interface{}) {
		return NewOrderFromContext(ctx)
	})




	basicOrderTraitMapping["http://schema.org/confirmationNumber"] = func(x *schema.OrderTrait, s []string) {
		x.ConfirmationNumber = s[0]
	}






	basicOrderTraitMapping["http://schema.org/paymentMethodId"] = func(x *schema.OrderTrait, s []string) {
		x.PaymentMethodId = s[0]
	}





	basicOrderTraitMapping["http://schema.org/discountCode"] = func(x *schema.OrderTrait, s []string) {
		x.DiscountCode = s[0]
	}


	basicOrderTraitMapping["http://schema.org/discountCurrency"] = func(x *schema.OrderTrait, s []string) {
		x.DiscountCurrency = s[0]
	}


	basicOrderTraitMapping["http://schema.org/isGift"] = func(x *schema.OrderTrait, s []string) {
		var err error
		x.IsGift, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}





	basicOrderTraitMapping["http://schema.org/orderNumber"] = func(x *schema.OrderTrait, s []string) {
		x.OrderNumber = s[0]
	}





	basicOrderTraitMapping["http://schema.org/paymentUrl"] = func(x *schema.OrderTrait, s []string) {
		x.PaymentUrl = s[0]
	}


	customOrderTraitMapping["http://schema.org/seller"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
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

	customOrderTraitMapping["http://schema.org/broker"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
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

	customOrderTraitMapping["http://schema.org/customer"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Customer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Customer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Customer = s
		}
	}

	customOrderTraitMapping["http://schema.org/paymentDue"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.PaymentDue = &y

		return
	}

	customOrderTraitMapping["http://schema.org/paymentDueDate"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.PaymentDueDate = &y

		return
	}

	customOrderTraitMapping["http://schema.org/paymentMethod"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.PaymentMethod
		if strings.HasPrefix(s, "_:") {
			y = NewPaymentMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPaymentMethod()
			y.Id = s
		}

		x.PaymentMethod = &y

		return
	}

	customOrderTraitMapping["http://schema.org/acceptedOffer"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.AcceptedOffer = &y

		return
	}

	customOrderTraitMapping["http://schema.org/billingAddress"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.PostalAddress
		if strings.HasPrefix(s, "_:") {
			y = NewPostalAddressFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPostalAddress()
			y.Id = s
		}

		x.BillingAddress = &y

		return
	}

	customOrderTraitMapping["http://schema.org/discount"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Discount, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Discount = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Discount = s
		}
	}

	customOrderTraitMapping["http://schema.org/merchant"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Merchant, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Merchant = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Merchant = s
		}
	}

	customOrderTraitMapping["http://schema.org/orderDate"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.OrderDate = &y

		return
	}

	customOrderTraitMapping["http://schema.org/orderDelivery"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.ParcelDelivery
		if strings.HasPrefix(s, "_:") {
			y = NewParcelDeliveryFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewParcelDelivery()
			y.Id = s
		}

		x.OrderDelivery = &y

		return
	}

	customOrderTraitMapping["http://schema.org/orderStatus"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.OrderStatus
		if strings.HasPrefix(s, "_:") {
			y = NewOrderStatusFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrderStatus()
			y.Id = s
		}

		x.OrderStatus = &y

		return
	}

	customOrderTraitMapping["http://schema.org/orderedItem"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.OrderedItem, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.OrderedItem = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.OrderedItem = s
		}
	}

	customOrderTraitMapping["http://schema.org/partOfInvoice"] = func(x *schema.OrderTrait, ctx jsonld.Context, s string){
		var y schema.Invoice
		if strings.HasPrefix(s, "_:") {
			y = NewInvoiceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInvoice()
			y.Id = s
		}

		x.PartOfInvoice = &y

		return
	}
}

func NewOrderFromContext(ctx jsonld.Context) (x schema.Order) {
	x.Type = "http://schema.org/Order"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOrderTrait(ctx, &x.OrderTrait)
	MapCustomToOrderTrait(ctx, &x.OrderTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOrderTrait(ctx jsonld.Context, x *schema.OrderTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOrderTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOrderTrait(ctx jsonld.Context, x *schema.OrderTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOrderTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}