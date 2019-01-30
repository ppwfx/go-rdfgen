package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsProduct strings.Replacer
var strconvProduct strconv.NumError

var basicProductTraitMapping = map[string]func(*schema.ProductTrait, []string){}
var customProductTraitMapping = map[string]func(*schema.ProductTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Product", func(ctx jsonld.Context) (interface{}) {
		return NewProductFromContext(ctx)
	})




	basicProductTraitMapping["http://schema.org/award"] = func(x *schema.ProductTrait, s []string) {
		x.Award = s[0]
	}


	basicProductTraitMapping["http://schema.org/awards"] = func(x *schema.ProductTrait, s []string) {
		x.Awards = s[0]
	}








	basicProductTraitMapping["http://schema.org/gtin12"] = func(x *schema.ProductTrait, s []string) {
		x.Gtin12 = s[0]
	}


	basicProductTraitMapping["http://schema.org/gtin13"] = func(x *schema.ProductTrait, s []string) {
		x.Gtin13 = s[0]
	}


	basicProductTraitMapping["http://schema.org/gtin14"] = func(x *schema.ProductTrait, s []string) {
		x.Gtin14 = s[0]
	}


	basicProductTraitMapping["http://schema.org/gtin8"] = func(x *schema.ProductTrait, s []string) {
		x.Gtin8 = s[0]
	}



	basicProductTraitMapping["http://schema.org/mpn"] = func(x *schema.ProductTrait, s []string) {
		x.Mpn = s[0]
	}


	basicProductTraitMapping["http://schema.org/sku"] = func(x *schema.ProductTrait, s []string) {
		x.Sku = s[0]
	}









	basicProductTraitMapping["http://schema.org/color"] = func(x *schema.ProductTrait, s []string) {
		x.Color = s[0]
	}







	basicProductTraitMapping["http://schema.org/productID"] = func(x *schema.ProductTrait, s []string) {
		x.ProductID = s[0]
	}





	customProductTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.AggregateRating
		if strings.HasPrefix(s, "_:") {
			y = NewAggregateRatingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAggregateRating()
			y.Id = s
		}

		x.AggregateRating = &y

		return
	}

	customProductTraitMapping["http://schema.org/audience"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Audience
		if strings.HasPrefix(s, "_:") {
			y = NewAudienceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAudience()
			y.Id = s
		}

		x.Audience = &y

		return
	}

	customProductTraitMapping["http://schema.org/height"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Height, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Height = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Height = s
		}
	}

	customProductTraitMapping["http://schema.org/material"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Material, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Material = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Material = s
		}
	}

	customProductTraitMapping["http://schema.org/offers"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
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

	customProductTraitMapping["http://schema.org/review"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.Review = &y

		return
	}

	customProductTraitMapping["http://schema.org/reviews"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.Reviews = &y

		return
	}

	customProductTraitMapping["http://schema.org/width"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Width, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Width = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Width = s
		}
	}

	customProductTraitMapping["http://schema.org/itemCondition"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.OfferItemCondition
		if strings.HasPrefix(s, "_:") {
			y = NewOfferItemConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOfferItemCondition()
			y.Id = s
		}

		x.ItemCondition = &y

		return
	}

	customProductTraitMapping["http://schema.org/category"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Category, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Category = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Category = s
		}
	}

	customProductTraitMapping["http://schema.org/additionalProperty"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.PropertyValue
		if strings.HasPrefix(s, "_:") {
			y = NewPropertyValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPropertyValue()
			y.Id = s
		}

		x.AdditionalProperty = &y

		return
	}

	customProductTraitMapping["http://schema.org/logo"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Logo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Logo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Logo = s
		}
	}

	customProductTraitMapping["http://schema.org/brand"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Brand, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Brand = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Brand = s
		}
	}

	customProductTraitMapping["http://schema.org/weight"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Weight = &y

		return
	}

	customProductTraitMapping["http://schema.org/isRelatedTo"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsRelatedTo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsRelatedTo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsRelatedTo = s
		}
	}

	customProductTraitMapping["http://schema.org/isSimilarTo"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsSimilarTo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsSimilarTo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsSimilarTo = s
		}
	}

	customProductTraitMapping["http://schema.org/depth"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Depth, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Depth = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Depth = s
		}
	}

	customProductTraitMapping["http://schema.org/isAccessoryOrSparePartFor"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Product
		if strings.HasPrefix(s, "_:") {
			y = NewProductFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProduct()
			y.Id = s
		}

		x.IsAccessoryOrSparePartFor = &y

		return
	}

	customProductTraitMapping["http://schema.org/isConsumableFor"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Product
		if strings.HasPrefix(s, "_:") {
			y = NewProductFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProduct()
			y.Id = s
		}

		x.IsConsumableFor = &y

		return
	}

	customProductTraitMapping["http://schema.org/manufacturer"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Manufacturer = &y

		return
	}

	customProductTraitMapping["http://schema.org/model"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Model, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Model = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Model = s
		}
	}

	customProductTraitMapping["http://schema.org/productionDate"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.ProductionDate = &y

		return
	}

	customProductTraitMapping["http://schema.org/purchaseDate"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.PurchaseDate = &y

		return
	}

	customProductTraitMapping["http://schema.org/releaseDate"] = func(x *schema.ProductTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.ReleaseDate = &y

		return
	}
}

func NewProductFromContext(ctx jsonld.Context) (x schema.Product) {
	x.Type = "http://schema.org/Product"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToProductTrait(ctx jsonld.Context, x *schema.ProductTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicProductTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToProductTrait(ctx jsonld.Context, x *schema.ProductTrait) () {
	for k, v := range ctx.Current {
		f, ok := customProductTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}