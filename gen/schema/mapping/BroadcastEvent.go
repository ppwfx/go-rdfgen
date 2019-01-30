package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBroadcastEvent strings.Replacer
var strconvBroadcastEvent strconv.NumError

var basicBroadcastEventTraitMapping = map[string]func(*schema.BroadcastEventTrait, []string){}
var customBroadcastEventTraitMapping = map[string]func(*schema.BroadcastEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BroadcastEvent", func(ctx jsonld.Context) (interface{}) {
		return NewBroadcastEventFromContext(ctx)
	})


	basicBroadcastEventTraitMapping["http://schema.org/videoFormat"] = func(x *schema.BroadcastEventTrait, s []string) {
		x.VideoFormat = s[0]
	}



	basicBroadcastEventTraitMapping["http://schema.org/isLiveBroadcast"] = func(x *schema.BroadcastEventTrait, s []string) {
		var err error
		x.IsLiveBroadcast, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	customBroadcastEventTraitMapping["http://schema.org/broadcastOfEvent"] = func(x *schema.BroadcastEventTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.BroadcastOfEvent = &y

		return
	}
}

func NewBroadcastEventFromContext(ctx jsonld.Context) (x schema.BroadcastEvent) {
	x.Type = "http://schema.org/BroadcastEvent"
	MapBasicToPublicationEventTrait(ctx, &x.PublicationEventTrait)
	MapCustomToPublicationEventTrait(ctx, &x.PublicationEventTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBroadcastEventTrait(ctx, &x.BroadcastEventTrait)
	MapCustomToBroadcastEventTrait(ctx, &x.BroadcastEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBroadcastEventTrait(ctx jsonld.Context, x *schema.BroadcastEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBroadcastEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBroadcastEventTrait(ctx jsonld.Context, x *schema.BroadcastEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBroadcastEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}