package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOrderItem strings.Replacer
var strconvOrderItem strconv.NumError

var basicOrderItemTraitMapping = map[string]func(*schema.OrderItemTrait, []string){}
var customOrderItemTraitMapping = map[string]func(*schema.OrderItemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OrderItem", func(ctx jsonld.Context) (interface{}) {
		return NewOrderItemFromContext(ctx)
	})




	basicOrderItemTraitMapping["http://schema.org/orderItemNumber"] = func(x *schema.OrderItemTrait, s []string) {
		x.OrderItemNumber = s[0]
	}



	basicOrderItemTraitMapping["http://schema.org/orderQuantity"] = func(x *schema.OrderItemTrait, s []string) {
		var err error
		x.OrderQuantity, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	customOrderItemTraitMapping["http://schema.org/orderDelivery"] = func(x *schema.OrderItemTrait, ctx jsonld.Context, s string){
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

	customOrderItemTraitMapping["http://schema.org/orderedItem"] = func(x *schema.OrderItemTrait, ctx jsonld.Context, s string){
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

	customOrderItemTraitMapping["http://schema.org/orderItemStatus"] = func(x *schema.OrderItemTrait, ctx jsonld.Context, s string){
		var y schema.OrderStatus
		if strings.HasPrefix(s, "_:") {
			y = NewOrderStatusFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrderStatus()
			y.Id = s
		}

		x.OrderItemStatus = &y

		return
	}
}

func NewOrderItemFromContext(ctx jsonld.Context) (x schema.OrderItem) {
	x.Type = "http://schema.org/OrderItem"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOrderItemTrait(ctx, &x.OrderItemTrait)
	MapCustomToOrderItemTrait(ctx, &x.OrderItemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOrderItemTrait(ctx jsonld.Context, x *schema.OrderItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOrderItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOrderItemTrait(ctx jsonld.Context, x *schema.OrderItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOrderItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}