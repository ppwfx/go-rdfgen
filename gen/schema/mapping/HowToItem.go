package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowToItem strings.Replacer
var strconvHowToItem strconv.NumError

var basicHowToItemTraitMapping = map[string]func(*schema.HowToItemTrait, []string){}
var customHowToItemTraitMapping = map[string]func(*schema.HowToItemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowToItem", func(ctx jsonld.Context) (interface{}) {
		return NewHowToItemFromContext(ctx)
	})



	customHowToItemTraitMapping["http://schema.org/requiredQuantity"] = func(x *schema.HowToItemTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RequiredQuantity, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RequiredQuantity = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RequiredQuantity = s
		}
	}
}

func NewHowToItemFromContext(ctx jsonld.Context) (x schema.HowToItem) {
	x.Type = "http://schema.org/HowToItem"
	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHowToItemTrait(ctx, &x.HowToItemTrait)
	MapCustomToHowToItemTrait(ctx, &x.HowToItemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToItemTrait(ctx jsonld.Context, x *schema.HowToItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToItemTrait(ctx jsonld.Context, x *schema.HowToItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}