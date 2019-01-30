package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsContactPoint strings.Replacer
var strconvContactPoint strconv.NumError

var basicContactPointTraitMapping = map[string]func(*schema.ContactPointTrait, []string){}
var customContactPointTraitMapping = map[string]func(*schema.ContactPointTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ContactPoint", func(ctx jsonld.Context) (interface{}) {
		return NewContactPointFromContext(ctx)
	})






	basicContactPointTraitMapping["http://schema.org/contactType"] = func(x *schema.ContactPointTrait, s []string) {
		x.ContactType = s[0]
	}


	basicContactPointTraitMapping["http://schema.org/email"] = func(x *schema.ContactPointTrait, s []string) {
		x.Email = s[0]
	}


	basicContactPointTraitMapping["http://schema.org/faxNumber"] = func(x *schema.ContactPointTrait, s []string) {
		x.FaxNumber = s[0]
	}




	basicContactPointTraitMapping["http://schema.org/telephone"] = func(x *schema.ContactPointTrait, s []string) {
		x.Telephone = s[0]
	}


	customContactPointTraitMapping["http://schema.org/hoursAvailable"] = func(x *schema.ContactPointTrait, ctx jsonld.Context, s string){
		var y schema.OpeningHoursSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewOpeningHoursSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOpeningHoursSpecification()
			y.Id = s
		}

		x.HoursAvailable = &y

		return
	}

	customContactPointTraitMapping["http://schema.org/areaServed"] = func(x *schema.ContactPointTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AreaServed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AreaServed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AreaServed = s
		}
	}

	customContactPointTraitMapping["http://schema.org/availableLanguage"] = func(x *schema.ContactPointTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AvailableLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AvailableLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AvailableLanguage = s
		}
	}

	customContactPointTraitMapping["http://schema.org/contactOption"] = func(x *schema.ContactPointTrait, ctx jsonld.Context, s string){
		var y schema.ContactPointOption
		if strings.HasPrefix(s, "_:") {
			y = NewContactPointOptionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewContactPointOption()
			y.Id = s
		}

		x.ContactOption = &y

		return
	}

	customContactPointTraitMapping["http://schema.org/productSupported"] = func(x *schema.ContactPointTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ProductSupported, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ProductSupported = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ProductSupported = s
		}
	}

	customContactPointTraitMapping["http://schema.org/serviceArea"] = func(x *schema.ContactPointTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ServiceArea, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ServiceArea = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ServiceArea = s
		}
	}
}

func NewContactPointFromContext(ctx jsonld.Context) (x schema.ContactPoint) {
	x.Type = "http://schema.org/ContactPoint"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToContactPointTrait(ctx, &x.ContactPointTrait)
	MapCustomToContactPointTrait(ctx, &x.ContactPointTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToContactPointTrait(ctx jsonld.Context, x *schema.ContactPointTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicContactPointTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToContactPointTrait(ctx jsonld.Context, x *schema.ContactPointTrait) () {
	for k, v := range ctx.Current {
		f, ok := customContactPointTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}