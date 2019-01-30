package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFoodEstablishment strings.Replacer
var strconvFoodEstablishment strconv.NumError

var basicFoodEstablishmentTraitMapping = map[string]func(*schema.FoodEstablishmentTrait, []string){}
var customFoodEstablishmentTraitMapping = map[string]func(*schema.FoodEstablishmentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FoodEstablishment", func(ctx jsonld.Context) (interface{}) {
		return NewFoodEstablishmentFromContext(ctx)
	})




	basicFoodEstablishmentTraitMapping["http://schema.org/servesCuisine"] = func(x *schema.FoodEstablishmentTrait, s []string) {
		x.ServesCuisine = s[0]
	}




	customFoodEstablishmentTraitMapping["http://schema.org/starRating"] = func(x *schema.FoodEstablishmentTrait, ctx jsonld.Context, s string){
		var y schema.Rating
		if strings.HasPrefix(s, "_:") {
			y = NewRatingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRating()
			y.Id = s
		}

		x.StarRating = &y

		return
	}

	customFoodEstablishmentTraitMapping["http://schema.org/hasMenu"] = func(x *schema.FoodEstablishmentTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.HasMenu, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.HasMenu = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.HasMenu = s
		}
	}

	customFoodEstablishmentTraitMapping["http://schema.org/acceptsReservations"] = func(x *schema.FoodEstablishmentTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AcceptsReservations, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AcceptsReservations = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AcceptsReservations = s
		}
	}

	customFoodEstablishmentTraitMapping["http://schema.org/menu"] = func(x *schema.FoodEstablishmentTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Menu, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Menu = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Menu = s
		}
	}
}

func NewFoodEstablishmentFromContext(ctx jsonld.Context) (x schema.FoodEstablishment) {
	x.Type = "http://schema.org/FoodEstablishment"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)
	MapCustomToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFoodEstablishmentTrait(ctx jsonld.Context, x *schema.FoodEstablishmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFoodEstablishmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFoodEstablishmentTrait(ctx jsonld.Context, x *schema.FoodEstablishmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFoodEstablishmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}