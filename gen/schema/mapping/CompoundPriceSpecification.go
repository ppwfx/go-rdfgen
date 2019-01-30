package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCompoundPriceSpecification strings.Replacer
var strconvCompoundPriceSpecification strconv.NumError

var basicCompoundPriceSpecificationTraitMapping = map[string]func(*schema.CompoundPriceSpecificationTrait, []string){}
var customCompoundPriceSpecificationTraitMapping = map[string]func(*schema.CompoundPriceSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CompoundPriceSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewCompoundPriceSpecificationFromContext(ctx)
	})



	customCompoundPriceSpecificationTraitMapping["http://schema.org/priceComponent"] = func(x *schema.CompoundPriceSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.UnitPriceSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewUnitPriceSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewUnitPriceSpecification()
			y.Id = s
		}

		x.PriceComponent = &y

		return
	}
}

func NewCompoundPriceSpecificationFromContext(ctx jsonld.Context) (x schema.CompoundPriceSpecification) {
	x.Type = "http://schema.org/CompoundPriceSpecification"
	MapBasicToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)
	MapCustomToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCompoundPriceSpecificationTrait(ctx, &x.CompoundPriceSpecificationTrait)
	MapCustomToCompoundPriceSpecificationTrait(ctx, &x.CompoundPriceSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCompoundPriceSpecificationTrait(ctx jsonld.Context, x *schema.CompoundPriceSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCompoundPriceSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCompoundPriceSpecificationTrait(ctx jsonld.Context, x *schema.CompoundPriceSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCompoundPriceSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}