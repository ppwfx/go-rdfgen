package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWarrantyPromise strings.Replacer
var strconvWarrantyPromise strconv.NumError

var basicWarrantyPromiseTraitMapping = map[string]func(*schema.WarrantyPromiseTrait, []string){}
var customWarrantyPromiseTraitMapping = map[string]func(*schema.WarrantyPromiseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WarrantyPromise", func(ctx jsonld.Context) (interface{}) {
		return NewWarrantyPromiseFromContext(ctx)
	})




	customWarrantyPromiseTraitMapping["http://schema.org/durationOfWarranty"] = func(x *schema.WarrantyPromiseTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.DurationOfWarranty = &y

		return
	}

	customWarrantyPromiseTraitMapping["http://schema.org/warrantyScope"] = func(x *schema.WarrantyPromiseTrait, ctx jsonld.Context, s string){
		var y schema.WarrantyScope
		if strings.HasPrefix(s, "_:") {
			y = NewWarrantyScopeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewWarrantyScope()
			y.Id = s
		}

		x.WarrantyScope = &y

		return
	}
}

func NewWarrantyPromiseFromContext(ctx jsonld.Context) (x schema.WarrantyPromise) {
	x.Type = "http://schema.org/WarrantyPromise"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWarrantyPromiseTrait(ctx, &x.WarrantyPromiseTrait)
	MapCustomToWarrantyPromiseTrait(ctx, &x.WarrantyPromiseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWarrantyPromiseTrait(ctx jsonld.Context, x *schema.WarrantyPromiseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWarrantyPromiseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWarrantyPromiseTrait(ctx jsonld.Context, x *schema.WarrantyPromiseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWarrantyPromiseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}