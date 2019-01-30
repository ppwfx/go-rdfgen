package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBroadcastChannel strings.Replacer
var strconvBroadcastChannel strconv.NumError

var basicBroadcastChannelTraitMapping = map[string]func(*schema.BroadcastChannelTrait, []string){}
var customBroadcastChannelTraitMapping = map[string]func(*schema.BroadcastChannelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BroadcastChannel", func(ctx jsonld.Context) (interface{}) {
		return NewBroadcastChannelFromContext(ctx)
	})



	basicBroadcastChannelTraitMapping["http://schema.org/broadcastChannelId"] = func(x *schema.BroadcastChannelTrait, s []string) {
		x.BroadcastChannelId = s[0]
	}



	basicBroadcastChannelTraitMapping["http://schema.org/broadcastServiceTier"] = func(x *schema.BroadcastChannelTrait, s []string) {
		x.BroadcastServiceTier = s[0]
	}




	customBroadcastChannelTraitMapping["http://schema.org/genre"] = func(x *schema.BroadcastChannelTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Genre, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Genre = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Genre = s
		}
	}

	customBroadcastChannelTraitMapping["http://schema.org/broadcastFrequency"] = func(x *schema.BroadcastChannelTrait, ctx jsonld.Context, s string){
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

	customBroadcastChannelTraitMapping["http://schema.org/inBroadcastLineup"] = func(x *schema.BroadcastChannelTrait, ctx jsonld.Context, s string){
		var y schema.CableOrSatelliteService
		if strings.HasPrefix(s, "_:") {
			y = NewCableOrSatelliteServiceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCableOrSatelliteService()
			y.Id = s
		}

		x.InBroadcastLineup = &y

		return
	}

	customBroadcastChannelTraitMapping["http://schema.org/providesBroadcastService"] = func(x *schema.BroadcastChannelTrait, ctx jsonld.Context, s string){
		var y schema.BroadcastService
		if strings.HasPrefix(s, "_:") {
			y = NewBroadcastServiceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBroadcastService()
			y.Id = s
		}

		x.ProvidesBroadcastService = &y

		return
	}
}

func NewBroadcastChannelFromContext(ctx jsonld.Context) (x schema.BroadcastChannel) {
	x.Type = "http://schema.org/BroadcastChannel"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)
	MapCustomToBroadcastChannelTrait(ctx, &x.BroadcastChannelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBroadcastChannelTrait(ctx jsonld.Context, x *schema.BroadcastChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBroadcastChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBroadcastChannelTrait(ctx jsonld.Context, x *schema.BroadcastChannelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBroadcastChannelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}