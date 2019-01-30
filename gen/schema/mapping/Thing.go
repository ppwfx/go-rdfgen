package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsThing strings.Replacer
var strconvThing strconv.NumError

var basicThingTraitMapping = map[string]func(*schema.ThingTrait, []string){}
var customThingTraitMapping = map[string]func(*schema.ThingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Thing", func(ctx jsonld.Context) (interface{}) {
		return NewThingFromContext(ctx)
	})


	basicThingTraitMapping["http://schema.org/additionalType"] = func(x *schema.ThingTrait, s []string) {
		x.AdditionalType = s[0]
	}


	basicThingTraitMapping["http://schema.org/alternateName"] = func(x *schema.ThingTrait, s []string) {
		x.AlternateName = s[0]
	}


	basicThingTraitMapping["http://schema.org/description"] = func(x *schema.ThingTrait, s []string) {
		x.Description = s[0]
	}


	basicThingTraitMapping["http://schema.org/disambiguatingDescription"] = func(x *schema.ThingTrait, s []string) {
		x.DisambiguatingDescription = s[0]
	}





	basicThingTraitMapping["http://schema.org/name"] = func(x *schema.ThingTrait, s []string) {
		x.Name = s[0]
	}



	basicThingTraitMapping["http://schema.org/sameAs"] = func(x *schema.ThingTrait, s []string) {
		x.SameAs = s[0]
	}



	basicThingTraitMapping["http://schema.org/url"] = func(x *schema.ThingTrait, s []string) {
		x.Url = s[0]
	}


	customThingTraitMapping["http://schema.org/identifier"] = func(x *schema.ThingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Identifier, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Identifier = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Identifier = s
		}
	}

	customThingTraitMapping["http://schema.org/image"] = func(x *schema.ThingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Image, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Image = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Image = s
		}
	}

	customThingTraitMapping["http://schema.org/mainEntityOfPage"] = func(x *schema.ThingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MainEntityOfPage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MainEntityOfPage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MainEntityOfPage = s
		}
	}

	customThingTraitMapping["http://schema.org/potentialAction"] = func(x *schema.ThingTrait, ctx jsonld.Context, s string){
		var y schema.Action
		if strings.HasPrefix(s, "_:") {
			y = NewActionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAction()
			y.Id = s
		}

		x.PotentialAction = &y

		return
	}

	customThingTraitMapping["http://schema.org/subjectOf"] = func(x *schema.ThingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SubjectOf, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SubjectOf = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SubjectOf = s
		}
	}
}

func NewThingFromContext(ctx jsonld.Context) (x schema.Thing) {
	x.Type = "http://schema.org/Thing"

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToThingTrait(ctx jsonld.Context, x *schema.ThingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicThingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToThingTrait(ctx jsonld.Context, x *schema.ThingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customThingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}