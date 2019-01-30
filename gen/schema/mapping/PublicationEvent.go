package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPublicationEvent strings.Replacer
var strconvPublicationEvent strconv.NumError

var basicPublicationEventTraitMapping = map[string]func(*schema.PublicationEventTrait, []string){}
var customPublicationEventTraitMapping = map[string]func(*schema.PublicationEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PublicationEvent", func(ctx jsonld.Context) (interface{}) {
		return NewPublicationEventFromContext(ctx)
	})


	basicPublicationEventTraitMapping["http://schema.org/isAccessibleForFree"] = func(x *schema.PublicationEventTrait, s []string) {
		var err error
		x.IsAccessibleForFree, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicPublicationEventTraitMapping["http://schema.org/free"] = func(x *schema.PublicationEventTrait, s []string) {
		var err error
		x.Free, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}




	customPublicationEventTraitMapping["http://schema.org/publishedBy"] = func(x *schema.PublicationEventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PublishedBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PublishedBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PublishedBy = s
		}
	}

	customPublicationEventTraitMapping["http://schema.org/publishedOn"] = func(x *schema.PublicationEventTrait, ctx jsonld.Context, s string){
		var y schema.BroadcastService
		if strings.HasPrefix(s, "_:") {
			y = NewBroadcastServiceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBroadcastService()
			y.Id = s
		}

		x.PublishedOn = &y

		return
	}
}

func NewPublicationEventFromContext(ctx jsonld.Context) (x schema.PublicationEvent) {
	x.Type = "http://schema.org/PublicationEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPublicationEventTrait(ctx, &x.PublicationEventTrait)
	MapCustomToPublicationEventTrait(ctx, &x.PublicationEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPublicationEventTrait(ctx jsonld.Context, x *schema.PublicationEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPublicationEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPublicationEventTrait(ctx jsonld.Context, x *schema.PublicationEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPublicationEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}