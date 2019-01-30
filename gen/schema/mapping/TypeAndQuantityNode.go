package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTypeAndQuantityNode strings.Replacer
var strconvTypeAndQuantityNode strconv.NumError

var basicTypeAndQuantityNodeTraitMapping = map[string]func(*schema.TypeAndQuantityNodeTrait, []string){}
var customTypeAndQuantityNodeTraitMapping = map[string]func(*schema.TypeAndQuantityNodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TypeAndQuantityNode", func(ctx jsonld.Context) (interface{}) {
		return NewTypeAndQuantityNodeFromContext(ctx)
	})




	basicTypeAndQuantityNodeTraitMapping["http://schema.org/unitText"] = func(x *schema.TypeAndQuantityNodeTrait, s []string) {
		x.UnitText = s[0]
	}


	basicTypeAndQuantityNodeTraitMapping["http://schema.org/amountOfThisGood"] = func(x *schema.TypeAndQuantityNodeTrait, s []string) {
		var err error
		x.AmountOfThisGood, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	customTypeAndQuantityNodeTraitMapping["http://schema.org/businessFunction"] = func(x *schema.TypeAndQuantityNodeTrait, ctx jsonld.Context, s string){
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

	customTypeAndQuantityNodeTraitMapping["http://schema.org/unitCode"] = func(x *schema.TypeAndQuantityNodeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.UnitCode, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.UnitCode = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.UnitCode = s
		}
	}

	customTypeAndQuantityNodeTraitMapping["http://schema.org/typeOfGood"] = func(x *schema.TypeAndQuantityNodeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TypeOfGood, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TypeOfGood = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TypeOfGood = s
		}
	}
}

func NewTypeAndQuantityNodeFromContext(ctx jsonld.Context) (x schema.TypeAndQuantityNode) {
	x.Type = "http://schema.org/TypeAndQuantityNode"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTypeAndQuantityNodeTrait(ctx, &x.TypeAndQuantityNodeTrait)
	MapCustomToTypeAndQuantityNodeTrait(ctx, &x.TypeAndQuantityNodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTypeAndQuantityNodeTrait(ctx jsonld.Context, x *schema.TypeAndQuantityNodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTypeAndQuantityNodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTypeAndQuantityNodeTrait(ctx jsonld.Context, x *schema.TypeAndQuantityNodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTypeAndQuantityNodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}