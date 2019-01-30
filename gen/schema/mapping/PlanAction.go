package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPlanAction strings.Replacer
var strconvPlanAction strconv.NumError

var basicPlanActionTraitMapping = map[string]func(*schema.PlanActionTrait, []string){}
var customPlanActionTraitMapping = map[string]func(*schema.PlanActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PlanAction", func(ctx jsonld.Context) (interface{}) {
		return NewPlanActionFromContext(ctx)
	})



	customPlanActionTraitMapping["http://schema.org/scheduledTime"] = func(x *schema.PlanActionTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ScheduledTime = &y

		return
	}
}

func NewPlanActionFromContext(ctx jsonld.Context) (x schema.PlanAction) {
	x.Type = "http://schema.org/PlanAction"
	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPlanActionTrait(ctx, &x.PlanActionTrait)
	MapCustomToPlanActionTrait(ctx, &x.PlanActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPlanActionTrait(ctx jsonld.Context, x *schema.PlanActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPlanActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPlanActionTrait(ctx jsonld.Context, x *schema.PlanActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPlanActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}