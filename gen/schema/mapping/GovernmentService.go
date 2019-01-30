package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGovernmentService strings.Replacer
var strconvGovernmentService strconv.NumError

var basicGovernmentServiceTraitMapping = map[string]func(*schema.GovernmentServiceTrait, []string){}
var customGovernmentServiceTraitMapping = map[string]func(*schema.GovernmentServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GovernmentService", func(ctx jsonld.Context) (interface{}) {
		return NewGovernmentServiceFromContext(ctx)
	})



	customGovernmentServiceTraitMapping["http://schema.org/serviceOperator"] = func(x *schema.GovernmentServiceTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.ServiceOperator = &y

		return
	}
}

func NewGovernmentServiceFromContext(ctx jsonld.Context) (x schema.GovernmentService) {
	x.Type = "http://schema.org/GovernmentService"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGovernmentServiceTrait(ctx, &x.GovernmentServiceTrait)
	MapCustomToGovernmentServiceTrait(ctx, &x.GovernmentServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGovernmentServiceTrait(ctx jsonld.Context, x *schema.GovernmentServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGovernmentServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGovernmentServiceTrait(ctx jsonld.Context, x *schema.GovernmentServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGovernmentServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}