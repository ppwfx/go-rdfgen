package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowToSupply strings.Replacer
var strconvHowToSupply strconv.NumError

var basicHowToSupplyTraitMapping = map[string]func(*schema.HowToSupplyTrait, []string){}
var customHowToSupplyTraitMapping = map[string]func(*schema.HowToSupplyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowToSupply", func(ctx jsonld.Context) (interface{}) {
		return NewHowToSupplyFromContext(ctx)
	})



	customHowToSupplyTraitMapping["http://schema.org/estimatedCost"] = func(x *schema.HowToSupplyTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EstimatedCost, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EstimatedCost = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EstimatedCost = s
		}
	}
}

func NewHowToSupplyFromContext(ctx jsonld.Context) (x schema.HowToSupply) {
	x.Type = "http://schema.org/HowToSupply"
	MapBasicToHowToItemTrait(ctx, &x.HowToItemTrait)
	MapCustomToHowToItemTrait(ctx, &x.HowToItemTrait)

	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHowToSupplyTrait(ctx, &x.HowToSupplyTrait)
	MapCustomToHowToSupplyTrait(ctx, &x.HowToSupplyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToSupplyTrait(ctx jsonld.Context, x *schema.HowToSupplyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToSupplyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToSupplyTrait(ctx jsonld.Context, x *schema.HowToSupplyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToSupplyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}