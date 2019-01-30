package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsServiceChannel strings.Replacer
var strconvServiceChannel strconv.NumError

var basicServiceChannelTraitMapping = map[string]func(*schema.ServiceChannelTrait, []string){}
var customServiceChannelTraitMapping = map[string]func(*schema.ServiceChannelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ServiceChannel", func(ctx jsonld.Context) (interface{}) {
		return NewServiceChannelFromContext(ctx)
	})









	basicServiceChannelTraitMapping["http://schema.org/serviceUrl"] = func(x *schema.ServiceChannelTrait, s []string) {
		x.ServiceUrl = s[0]
	}


	customServiceChannelTraitMapping["http://schema.org/availableLanguage"] = func(x *schema.ServiceChannelTrait, ctx jsonld.Context, s string){
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

	customServiceChannelTraitMapping["http://schema.org/processingTime"] = func(x *schema.ServiceChannelTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.ProcessingTime = &y

		return
	}

	customServiceChannelTraitMapping["http://schema.org/providesService"] = func(x *schema.ServiceChannelTrait, ctx jsonld.Context, s string){
		var y schema.Service
		if strings.HasPrefix(s, "_:") {
			y = NewServiceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewService()
			y.Id = s
		}

		x.ProvidesService = &y

		return
	}

	customServiceChannelTraitMapping["http://schema.org/serviceLocation"] = func(x *schema.ServiceChannelTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ServiceLocation = &y

		return
	}

	customServiceChannelTraitMapping["http://schema.org/servicePhone"] = func(x *schema.ServiceChannelTrait, ctx jsonld.Context, s string){
		var y schema.ContactPoint
		if strings.HasPrefix(s, "_:") {
			y = NewContactPointFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewContactPoint()
			y.Id = s
		}

		x.ServicePhone = &y

		return
	}

	customServiceChannelTraitMapping["http://schema.org/servicePostalAddress"] = func(x *schema.ServiceChannelTrait, ctx jsonld.Context, s string){
		var y schema.PostalAddress
		if strings.HasPrefix(s, "_:") {
			y = NewPostalAddressFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPostalAddress()
			y.Id = s
		}

		x.ServicePostalAddress = &y

		return
	}

	customServiceChannelTraitMapping["http://schema.org/serviceSmsNumber"] = func(x *schema.ServiceChannelTrait, ctx jsonld.Context, s string){
		var y schema.ContactPoint
		if strings.HasPrefix(s, "_:") {
			y = NewContactPointFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewContactPoint()
			y.Id = s
		}

		x.ServiceSmsNumber = &y

		return
	}
}

func NewServiceChannelFromContext(ctx jsonld.Context) (x schema.ServiceChannel) {
	x.Type = "http://schema.org/ServiceChannel"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToServiceChannelTrait(ctx, &x.ServiceChannelTrait)
	MapCustomToServiceChannelTrait(ctx, &x.ServiceChannelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToServiceChannelTrait(ctx jsonld.Context, x *schema.ServiceChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicServiceChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToServiceChannelTrait(ctx jsonld.Context, x *schema.ServiceChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customServiceChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}