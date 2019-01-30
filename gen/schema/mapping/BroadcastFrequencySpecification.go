package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBroadcastFrequencySpecification strings.Replacer
var strconvBroadcastFrequencySpecification strconv.NumError

var basicBroadcastFrequencySpecificationTraitMapping = map[string]func(*schema.BroadcastFrequencySpecificationTrait, []string){}
var customBroadcastFrequencySpecificationTraitMapping = map[string]func(*schema.BroadcastFrequencySpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BroadcastFrequencySpecification", func(ctx jsonld.Context) (interface{}) {
		return NewBroadcastFrequencySpecificationFromContext(ctx)
	})



	customBroadcastFrequencySpecificationTraitMapping["http://schema.org/broadcastFrequencyValue"] = func(x *schema.BroadcastFrequencySpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BroadcastFrequencyValue, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BroadcastFrequencyValue = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BroadcastFrequencyValue = s
		}
	}
}

func NewBroadcastFrequencySpecificationFromContext(ctx jsonld.Context) (x schema.BroadcastFrequencySpecification) {
	x.Type = "http://schema.org/BroadcastFrequencySpecification"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBroadcastFrequencySpecificationTrait(ctx, &x.BroadcastFrequencySpecificationTrait)
	MapCustomToBroadcastFrequencySpecificationTrait(ctx, &x.BroadcastFrequencySpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBroadcastFrequencySpecificationTrait(ctx jsonld.Context, x *schema.BroadcastFrequencySpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBroadcastFrequencySpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBroadcastFrequencySpecificationTrait(ctx jsonld.Context, x *schema.BroadcastFrequencySpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBroadcastFrequencySpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}