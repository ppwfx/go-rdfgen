package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsActionAccessSpecification strings.Replacer
var strconvActionAccessSpecification strconv.NumError

var basicActionAccessSpecificationTraitMapping = map[string]func(*schema.ActionAccessSpecificationTrait, []string){}
var customActionAccessSpecificationTraitMapping = map[string]func(*schema.ActionAccessSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ActionAccessSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewActionAccessSpecificationFromContext(ctx)
	})








	customActionAccessSpecificationTraitMapping["http://schema.org/requiresSubscription"] = func(x *schema.ActionAccessSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RequiresSubscription, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RequiresSubscription = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RequiresSubscription = s
		}
	}

	customActionAccessSpecificationTraitMapping["http://schema.org/availabilityEnds"] = func(x *schema.ActionAccessSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.AvailabilityEnds = &y

		return
	}

	customActionAccessSpecificationTraitMapping["http://schema.org/availabilityStarts"] = func(x *schema.ActionAccessSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.AvailabilityStarts = &y

		return
	}

	customActionAccessSpecificationTraitMapping["http://schema.org/eligibleRegion"] = func(x *schema.ActionAccessSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EligibleRegion, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EligibleRegion = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EligibleRegion = s
		}
	}

	customActionAccessSpecificationTraitMapping["http://schema.org/category"] = func(x *schema.ActionAccessSpecificationTrait, ctx jsonld.Context, s string){
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

	customActionAccessSpecificationTraitMapping["http://schema.org/expectsAcceptanceOf"] = func(x *schema.ActionAccessSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.ExpectsAcceptanceOf = &y

		return
	}
}

func NewActionAccessSpecificationFromContext(ctx jsonld.Context) (x schema.ActionAccessSpecification) {
	x.Type = "http://schema.org/ActionAccessSpecification"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToActionAccessSpecificationTrait(ctx, &x.ActionAccessSpecificationTrait)
	MapCustomToActionAccessSpecificationTrait(ctx, &x.ActionAccessSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToActionAccessSpecificationTrait(ctx jsonld.Context, x *schema.ActionAccessSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicActionAccessSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToActionAccessSpecificationTrait(ctx jsonld.Context, x *schema.ActionAccessSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customActionAccessSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}