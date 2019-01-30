package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSportsTeam strings.Replacer
var strconvSportsTeam strconv.NumError

var basicSportsTeamTraitMapping = map[string]func(*schema.SportsTeamTrait, []string){}
var customSportsTeamTraitMapping = map[string]func(*schema.SportsTeamTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SportsTeam", func(ctx jsonld.Context) (interface{}) {
		return NewSportsTeamFromContext(ctx)
	})




	customSportsTeamTraitMapping["http://schema.org/coach"] = func(x *schema.SportsTeamTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Coach = &y

		return
	}

	customSportsTeamTraitMapping["http://schema.org/athlete"] = func(x *schema.SportsTeamTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Athlete = &y

		return
	}
}

func NewSportsTeamFromContext(ctx jsonld.Context) (x schema.SportsTeam) {
	x.Type = "http://schema.org/SportsTeam"
	MapBasicToSportsOrganizationTrait(ctx, &x.SportsOrganizationTrait)
	MapCustomToSportsOrganizationTrait(ctx, &x.SportsOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSportsTeamTrait(ctx, &x.SportsTeamTrait)
	MapCustomToSportsTeamTrait(ctx, &x.SportsTeamTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSportsTeamTrait(ctx jsonld.Context, x *schema.SportsTeamTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSportsTeamTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSportsTeamTrait(ctx jsonld.Context, x *schema.SportsTeamTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSportsTeamTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}