package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserPageVisits strings.Replacer
var strconvUserPageVisits strconv.NumError

var basicUserPageVisitsTraitMapping = map[string]func(*schema.UserPageVisitsTrait, []string){}
var customUserPageVisitsTraitMapping = map[string]func(*schema.UserPageVisitsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserPageVisits", func(ctx jsonld.Context) (interface{}) {
		return NewUserPageVisitsFromContext(ctx)
	})

}

func NewUserPageVisitsFromContext(ctx jsonld.Context) (x schema.UserPageVisits) {
	x.Type = "http://schema.org/UserPageVisits"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserPageVisitsTrait(ctx, &x.UserPageVisitsTrait)
	MapCustomToUserPageVisitsTrait(ctx, &x.UserPageVisitsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserPageVisitsTrait(ctx jsonld.Context, x *schema.UserPageVisitsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserPageVisitsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserPageVisitsTrait(ctx jsonld.Context, x *schema.UserPageVisitsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserPageVisitsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}