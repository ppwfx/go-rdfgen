package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBroadcastService strings.Replacer
var strconvBroadcastService strconv.NumError

var basicBroadcastServiceTraitMapping = map[string]func(*schema.BroadcastServiceTrait, []string){}
var customBroadcastServiceTraitMapping = map[string]func(*schema.BroadcastServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BroadcastService", func(ctx jsonld.Context) (interface{}) {
		return NewBroadcastServiceFromContext(ctx)
	})



	basicBroadcastServiceTraitMapping["http://schema.org/videoFormat"] = func(x *schema.BroadcastServiceTrait, s []string) {
		x.VideoFormat = s[0]
	}


	basicBroadcastServiceTraitMapping["http://schema.org/broadcastDisplayName"] = func(x *schema.BroadcastServiceTrait, s []string) {
		x.BroadcastDisplayName = s[0]
	}


	basicBroadcastServiceTraitMapping["http://schema.org/broadcastTimezone"] = func(x *schema.BroadcastServiceTrait, s []string) {
		x.BroadcastTimezone = s[0]
	}







	customBroadcastServiceTraitMapping["http://schema.org/broadcastFrequency"] = func(x *schema.BroadcastServiceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BroadcastFrequency, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BroadcastFrequency = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BroadcastFrequency = s
		}
	}

	customBroadcastServiceTraitMapping["http://schema.org/area"] = func(x *schema.BroadcastServiceTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.Area = &y

		return
	}

	customBroadcastServiceTraitMapping["http://schema.org/hasBroadcastChannel"] = func(x *schema.BroadcastServiceTrait, ctx jsonld.Context, s string){
		var y schema.BroadcastChannel
		if strings.HasPrefix(s, "_:") {
			y = NewBroadcastChannelFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBroadcastChannel()
			y.Id = s
		}

		x.HasBroadcastChannel = &y

		return
	}

	customBroadcastServiceTraitMapping["http://schema.org/broadcastAffiliateOf"] = func(x *schema.BroadcastServiceTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.BroadcastAffiliateOf = &y

		return
	}

	customBroadcastServiceTraitMapping["http://schema.org/broadcaster"] = func(x *schema.BroadcastServiceTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Broadcaster = &y

		return
	}

	customBroadcastServiceTraitMapping["http://schema.org/parentService"] = func(x *schema.BroadcastServiceTrait, ctx jsonld.Context, s string){
		var y schema.BroadcastService
		if strings.HasPrefix(s, "_:") {
			y = NewBroadcastServiceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBroadcastService()
			y.Id = s
		}

		x.ParentService = &y

		return
	}
}

func NewBroadcastServiceFromContext(ctx jsonld.Context) (x schema.BroadcastService) {
	x.Type = "http://schema.org/BroadcastService"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBroadcastServiceTrait(ctx, &x.BroadcastServiceTrait)
	MapCustomToBroadcastServiceTrait(ctx, &x.BroadcastServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBroadcastServiceTrait(ctx jsonld.Context, x *schema.BroadcastServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBroadcastServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBroadcastServiceTrait(ctx jsonld.Context, x *schema.BroadcastServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBroadcastServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}