package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOnDemandEvent strings.Replacer
var strconvOnDemandEvent strconv.NumError

var basicOnDemandEventTraitMapping = map[string]func(*schema.OnDemandEventTrait, []string){}
var customOnDemandEventTraitMapping = map[string]func(*schema.OnDemandEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OnDemandEvent", func(ctx jsonld.Context) (interface{}) {
		return NewOnDemandEventFromContext(ctx)
	})

}

func NewOnDemandEventFromContext(ctx jsonld.Context) (x schema.OnDemandEvent) {
	x.Type = "http://schema.org/OnDemandEvent"
	MapBasicToPublicationEventTrait(ctx, &x.PublicationEventTrait)
	MapCustomToPublicationEventTrait(ctx, &x.PublicationEventTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOnDemandEventTrait(ctx, &x.OnDemandEventTrait)
	MapCustomToOnDemandEventTrait(ctx, &x.OnDemandEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOnDemandEventTrait(ctx jsonld.Context, x *schema.OnDemandEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOnDemandEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOnDemandEventTrait(ctx jsonld.Context, x *schema.OnDemandEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOnDemandEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}