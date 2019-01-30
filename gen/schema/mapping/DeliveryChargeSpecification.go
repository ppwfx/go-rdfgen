package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDeliveryChargeSpecification strings.Replacer
var strconvDeliveryChargeSpecification strconv.NumError

var basicDeliveryChargeSpecificationTraitMapping = map[string]func(*schema.DeliveryChargeSpecificationTrait, []string){}
var customDeliveryChargeSpecificationTraitMapping = map[string]func(*schema.DeliveryChargeSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DeliveryChargeSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewDeliveryChargeSpecificationFromContext(ctx)
	})






	customDeliveryChargeSpecificationTraitMapping["http://schema.org/areaServed"] = func(x *schema.DeliveryChargeSpecificationTrait, ctx jsonld.Context, s string){
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

	customDeliveryChargeSpecificationTraitMapping["http://schema.org/eligibleRegion"] = func(x *schema.DeliveryChargeSpecificationTrait, ctx jsonld.Context, s string){
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

	customDeliveryChargeSpecificationTraitMapping["http://schema.org/ineligibleRegion"] = func(x *schema.DeliveryChargeSpecificationTrait, ctx jsonld.Context, s string){
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

	customDeliveryChargeSpecificationTraitMapping["http://schema.org/appliesToDeliveryMethod"] = func(x *schema.DeliveryChargeSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DeliveryMethod
		if strings.HasPrefix(s, "_:") {
			y = NewDeliveryMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDeliveryMethod()
			y.Id = s
		}

		x.AppliesToDeliveryMethod = &y

		return
	}
}

func NewDeliveryChargeSpecificationFromContext(ctx jsonld.Context) (x schema.DeliveryChargeSpecification) {
	x.Type = "http://schema.org/DeliveryChargeSpecification"
	MapBasicToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)
	MapCustomToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDeliveryChargeSpecificationTrait(ctx, &x.DeliveryChargeSpecificationTrait)
	MapCustomToDeliveryChargeSpecificationTrait(ctx, &x.DeliveryChargeSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDeliveryChargeSpecificationTrait(ctx jsonld.Context, x *schema.DeliveryChargeSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDeliveryChargeSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDeliveryChargeSpecificationTrait(ctx jsonld.Context, x *schema.DeliveryChargeSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDeliveryChargeSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}