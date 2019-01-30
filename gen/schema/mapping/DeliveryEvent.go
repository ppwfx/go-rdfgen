package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDeliveryEvent strings.Replacer
var strconvDeliveryEvent strconv.NumError

var basicDeliveryEventTraitMapping = map[string]func(*schema.DeliveryEventTrait, []string){}
var customDeliveryEventTraitMapping = map[string]func(*schema.DeliveryEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DeliveryEvent", func(ctx jsonld.Context) (interface{}) {
		return NewDeliveryEventFromContext(ctx)
	})





	basicDeliveryEventTraitMapping["http://schema.org/accessCode"] = func(x *schema.DeliveryEventTrait, s []string) {
		x.AccessCode = s[0]
	}


	customDeliveryEventTraitMapping["http://schema.org/hasDeliveryMethod"] = func(x *schema.DeliveryEventTrait, ctx jsonld.Context, s string){
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

	customDeliveryEventTraitMapping["http://schema.org/availableFrom"] = func(x *schema.DeliveryEventTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.AvailableFrom = &y

		return
	}

	customDeliveryEventTraitMapping["http://schema.org/availableThrough"] = func(x *schema.DeliveryEventTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.AvailableThrough = &y

		return
	}
}

func NewDeliveryEventFromContext(ctx jsonld.Context) (x schema.DeliveryEvent) {
	x.Type = "http://schema.org/DeliveryEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDeliveryEventTrait(ctx, &x.DeliveryEventTrait)
	MapCustomToDeliveryEventTrait(ctx, &x.DeliveryEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDeliveryEventTrait(ctx jsonld.Context, x *schema.DeliveryEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDeliveryEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDeliveryEventTrait(ctx jsonld.Context, x *schema.DeliveryEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDeliveryEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}