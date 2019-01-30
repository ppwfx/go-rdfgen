package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsItemList strings.Replacer
var strconvItemList strconv.NumError

var basicItemListTraitMapping = map[string]func(*schema.ItemListTrait, []string){}
var customItemListTraitMapping = map[string]func(*schema.ItemListTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ItemList", func(ctx jsonld.Context) (interface{}) {
		return NewItemListFromContext(ctx)
	})





	customItemListTraitMapping["http://schema.org/itemListElement"] = func(x *schema.ItemListTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ItemListElement, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ItemListElement = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ItemListElement = s
		}
	}

	customItemListTraitMapping["http://schema.org/itemListOrder"] = func(x *schema.ItemListTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ItemListOrder, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ItemListOrder = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ItemListOrder = s
		}
	}

	customItemListTraitMapping["http://schema.org/numberOfItems"] = func(x *schema.ItemListTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.NumberOfItems = &y

		return
	}
}

func NewItemListFromContext(ctx jsonld.Context) (x schema.ItemList) {
	x.Type = "http://schema.org/ItemList"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToItemListTrait(ctx, &x.ItemListTrait)
	MapCustomToItemListTrait(ctx, &x.ItemListTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToItemListTrait(ctx jsonld.Context, x *schema.ItemListTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicItemListTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToItemListTrait(ctx jsonld.Context, x *schema.ItemListTrait) () {
	for k, v := range ctx.Current {
		f, ok := customItemListTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}