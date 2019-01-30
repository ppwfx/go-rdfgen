package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMenuItem strings.Replacer
var strconvMenuItem strconv.NumError

var basicMenuItemTraitMapping = map[string]func(*schema.MenuItemTrait, []string){}
var customMenuItemTraitMapping = map[string]func(*schema.MenuItemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MenuItem", func(ctx jsonld.Context) (interface{}) {
		return NewMenuItemFromContext(ctx)
	})






	customMenuItemTraitMapping["http://schema.org/offers"] = func(x *schema.MenuItemTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.Offers = &y

		return
	}

	customMenuItemTraitMapping["http://schema.org/menuAddOn"] = func(x *schema.MenuItemTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MenuAddOn, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MenuAddOn = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MenuAddOn = s
		}
	}

	customMenuItemTraitMapping["http://schema.org/nutrition"] = func(x *schema.MenuItemTrait, ctx jsonld.Context, s string){
		var y schema.NutritionInformation
		if strings.HasPrefix(s, "_:") {
			y = NewNutritionInformationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewNutritionInformation()
			y.Id = s
		}

		x.Nutrition = &y

		return
	}

	customMenuItemTraitMapping["http://schema.org/suitableForDiet"] = func(x *schema.MenuItemTrait, ctx jsonld.Context, s string){
		var y schema.RestrictedDiet
		if strings.HasPrefix(s, "_:") {
			y = NewRestrictedDietFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRestrictedDiet()
			y.Id = s
		}

		x.SuitableForDiet = &y

		return
	}
}

func NewMenuItemFromContext(ctx jsonld.Context) (x schema.MenuItem) {
	x.Type = "http://schema.org/MenuItem"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMenuItemTrait(ctx, &x.MenuItemTrait)
	MapCustomToMenuItemTrait(ctx, &x.MenuItemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMenuItemTrait(ctx jsonld.Context, x *schema.MenuItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMenuItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMenuItemTrait(ctx jsonld.Context, x *schema.MenuItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMenuItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}