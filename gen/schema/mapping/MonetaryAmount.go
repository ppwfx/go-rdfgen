package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMonetaryAmount strings.Replacer
var strconvMonetaryAmount strconv.NumError

var basicMonetaryAmountTraitMapping = map[string]func(*schema.MonetaryAmountTrait, []string){}
var customMonetaryAmountTraitMapping = map[string]func(*schema.MonetaryAmountTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MonetaryAmount", func(ctx jsonld.Context) (interface{}) {
		return NewMonetaryAmountFromContext(ctx)
	})




	basicMonetaryAmountTraitMapping["http://schema.org/currency"] = func(x *schema.MonetaryAmountTrait, s []string) {
		x.Currency = s[0]
	}


	basicMonetaryAmountTraitMapping["http://schema.org/maxValue"] = func(x *schema.MonetaryAmountTrait, s []string) {
		var err error
		x.MaxValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicMonetaryAmountTraitMapping["http://schema.org/minValue"] = func(x *schema.MonetaryAmountTrait, s []string) {
		var err error
		x.MinValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	customMonetaryAmountTraitMapping["http://schema.org/validFrom"] = func(x *schema.MonetaryAmountTrait, ctx jsonld.Context, s string){
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

	customMonetaryAmountTraitMapping["http://schema.org/validThrough"] = func(x *schema.MonetaryAmountTrait, ctx jsonld.Context, s string){
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

	customMonetaryAmountTraitMapping["http://schema.org/value"] = func(x *schema.MonetaryAmountTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Value, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Value = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Value = s
		}
	}
}

func NewMonetaryAmountFromContext(ctx jsonld.Context) (x schema.MonetaryAmount) {
	x.Type = "http://schema.org/MonetaryAmount"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMonetaryAmountTrait(ctx, &x.MonetaryAmountTrait)
	MapCustomToMonetaryAmountTrait(ctx, &x.MonetaryAmountTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMonetaryAmountTrait(ctx jsonld.Context, x *schema.MonetaryAmountTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMonetaryAmountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMonetaryAmountTrait(ctx jsonld.Context, x *schema.MonetaryAmountTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMonetaryAmountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}