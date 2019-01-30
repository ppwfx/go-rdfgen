package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsScheduleAction strings.Replacer
var strconvScheduleAction strconv.NumError

var basicScheduleActionTraitMapping = map[string]func(*schema.ScheduleActionTrait, []string){}
var customScheduleActionTraitMapping = map[string]func(*schema.ScheduleActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ScheduleAction", func(ctx jsonld.Context) (interface{}) {
		return NewScheduleActionFromContext(ctx)
	})

}

func NewScheduleActionFromContext(ctx jsonld.Context) (x schema.ScheduleAction) {
	x.Type = "http://schema.org/ScheduleAction"
	MapBasicToPlanActionTrait(ctx, &x.PlanActionTrait)
	MapCustomToPlanActionTrait(ctx, &x.PlanActionTrait)

	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToScheduleActionTrait(ctx, &x.ScheduleActionTrait)
	MapCustomToScheduleActionTrait(ctx, &x.ScheduleActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToScheduleActionTrait(ctx jsonld.Context, x *schema.ScheduleActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicScheduleActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToScheduleActionTrait(ctx jsonld.Context, x *schema.ScheduleActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customScheduleActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}