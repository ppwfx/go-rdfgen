package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSportsOrganization strings.Replacer
var strconvSportsOrganization strconv.NumError

var basicSportsOrganizationTraitMapping = map[string]func(*schema.SportsOrganizationTrait, []string){}
var customSportsOrganizationTraitMapping = map[string]func(*schema.SportsOrganizationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SportsOrganization", func(ctx jsonld.Context) (interface{}) {
		return NewSportsOrganizationFromContext(ctx)
	})



	customSportsOrganizationTraitMapping["http://schema.org/sport"] = func(x *schema.SportsOrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Sport, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Sport = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Sport = s
		}
	}
}

func NewSportsOrganizationFromContext(ctx jsonld.Context) (x schema.SportsOrganization) {
	x.Type = "http://schema.org/SportsOrganization"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSportsOrganizationTrait(ctx, &x.SportsOrganizationTrait)
	MapCustomToSportsOrganizationTrait(ctx, &x.SportsOrganizationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSportsOrganizationTrait(ctx jsonld.Context, x *schema.SportsOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSportsOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSportsOrganizationTrait(ctx jsonld.Context, x *schema.SportsOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSportsOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}