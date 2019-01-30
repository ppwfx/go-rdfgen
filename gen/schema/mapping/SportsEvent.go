package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSportsEvent strings.Replacer
var strconvSportsEvent strconv.NumError

var basicSportsEventTraitMapping = map[string]func(*schema.SportsEventTrait, []string){}
var customSportsEventTraitMapping = map[string]func(*schema.SportsEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SportsEvent", func(ctx jsonld.Context) (interface{}) {
		return NewSportsEventFromContext(ctx)
	})





	customSportsEventTraitMapping["http://schema.org/awayTeam"] = func(x *schema.SportsEventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AwayTeam, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AwayTeam = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AwayTeam = s
		}
	}

	customSportsEventTraitMapping["http://schema.org/competitor"] = func(x *schema.SportsEventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Competitor, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Competitor = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Competitor = s
		}
	}

	customSportsEventTraitMapping["http://schema.org/homeTeam"] = func(x *schema.SportsEventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.HomeTeam, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.HomeTeam = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.HomeTeam = s
		}
	}
}

func NewSportsEventFromContext(ctx jsonld.Context) (x schema.SportsEvent) {
	x.Type = "http://schema.org/SportsEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSportsEventTrait(ctx, &x.SportsEventTrait)
	MapCustomToSportsEventTrait(ctx, &x.SportsEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSportsEventTrait(ctx jsonld.Context, x *schema.SportsEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSportsEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSportsEventTrait(ctx jsonld.Context, x *schema.SportsEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSportsEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}