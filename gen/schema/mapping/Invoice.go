package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInvoice strings.Replacer
var strconvInvoice strconv.NumError

var basicInvoiceTraitMapping = map[string]func(*schema.InvoiceTrait, []string){}
var customInvoiceTraitMapping = map[string]func(*schema.InvoiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Invoice", func(ctx jsonld.Context) (interface{}) {
		return NewInvoiceFromContext(ctx)
	})





	basicInvoiceTraitMapping["http://schema.org/accountId"] = func(x *schema.InvoiceTrait, s []string) {
		x.AccountId = s[0]
	}



	basicInvoiceTraitMapping["http://schema.org/confirmationNumber"] = func(x *schema.InvoiceTrait, s []string) {
		x.ConfirmationNumber = s[0]
	}







	basicInvoiceTraitMapping["http://schema.org/paymentMethodId"] = func(x *schema.InvoiceTrait, s []string) {
		x.PaymentMethodId = s[0]
	}






	customInvoiceTraitMapping["http://schema.org/provider"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
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

	customInvoiceTraitMapping["http://schema.org/category"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
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

	customInvoiceTraitMapping["http://schema.org/broker"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
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

	customInvoiceTraitMapping["http://schema.org/billingPeriod"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.BillingPeriod = &y

		return
	}

	customInvoiceTraitMapping["http://schema.org/customer"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
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

	customInvoiceTraitMapping["http://schema.org/minimumPaymentDue"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MinimumPaymentDue, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MinimumPaymentDue = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MinimumPaymentDue = s
		}
	}

	customInvoiceTraitMapping["http://schema.org/paymentDue"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
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

	customInvoiceTraitMapping["http://schema.org/paymentDueDate"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
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

	customInvoiceTraitMapping["http://schema.org/paymentMethod"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
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

	customInvoiceTraitMapping["http://schema.org/paymentStatus"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PaymentStatus, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PaymentStatus = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PaymentStatus = s
		}
	}

	customInvoiceTraitMapping["http://schema.org/referencesOrder"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
		var y schema.Order
		if strings.HasPrefix(s, "_:") {
			y = NewOrderFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrder()
			y.Id = s
		}

		x.ReferencesOrder = &y

		return
	}

	customInvoiceTraitMapping["http://schema.org/scheduledPaymentDate"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.ScheduledPaymentDate = &y

		return
	}

	customInvoiceTraitMapping["http://schema.org/totalPaymentDue"] = func(x *schema.InvoiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TotalPaymentDue, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TotalPaymentDue = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TotalPaymentDue = s
		}
	}
}

func NewInvoiceFromContext(ctx jsonld.Context) (x schema.Invoice) {
	x.Type = "http://schema.org/Invoice"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInvoiceTrait(ctx, &x.InvoiceTrait)
	MapCustomToInvoiceTrait(ctx, &x.InvoiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInvoiceTrait(ctx jsonld.Context, x *schema.InvoiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInvoiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInvoiceTrait(ctx jsonld.Context, x *schema.InvoiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInvoiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}