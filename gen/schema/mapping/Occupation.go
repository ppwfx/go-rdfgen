package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOccupation strings.Replacer
var strconvOccupation strconv.NumError

var basicOccupationTraitMapping = map[string]func(*schema.OccupationTrait, []string){}
var customOccupationTraitMapping = map[string]func(*schema.OccupationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Occupation", func(ctx jsonld.Context) (interface{}) {
		return NewOccupationFromContext(ctx)
	})


	basicOccupationTraitMapping["http://schema.org/educationRequirements"] = func(x *schema.OccupationTrait, s []string) {
		x.EducationRequirements = s[0]
	}



	basicOccupationTraitMapping["http://schema.org/experienceRequirements"] = func(x *schema.OccupationTrait, s []string) {
		x.ExperienceRequirements = s[0]
	}


	basicOccupationTraitMapping["http://schema.org/occupationalCategory"] = func(x *schema.OccupationTrait, s []string) {
		x.OccupationalCategory = s[0]
	}


	basicOccupationTraitMapping["http://schema.org/qualifications"] = func(x *schema.OccupationTrait, s []string) {
		x.Qualifications = s[0]
	}


	basicOccupationTraitMapping["http://schema.org/responsibilities"] = func(x *schema.OccupationTrait, s []string) {
		x.Responsibilities = s[0]
	}


	basicOccupationTraitMapping["http://schema.org/skills"] = func(x *schema.OccupationTrait, s []string) {
		x.Skills = s[0]
	}



	customOccupationTraitMapping["http://schema.org/estimatedSalary"] = func(x *schema.OccupationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EstimatedSalary, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EstimatedSalary = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EstimatedSalary = s
		}
	}

	customOccupationTraitMapping["http://schema.org/occupationLocation"] = func(x *schema.OccupationTrait, ctx jsonld.Context, s string){
		var y schema.AdministrativeArea
		if strings.HasPrefix(s, "_:") {
			y = NewAdministrativeAreaFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAdministrativeArea()
			y.Id = s
		}

		x.OccupationLocation = &y

		return
	}
}

func NewOccupationFromContext(ctx jsonld.Context) (x schema.Occupation) {
	x.Type = "http://schema.org/Occupation"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOccupationTrait(ctx, &x.OccupationTrait)
	MapCustomToOccupationTrait(ctx, &x.OccupationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOccupationTrait(ctx jsonld.Context, x *schema.OccupationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOccupationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOccupationTrait(ctx jsonld.Context, x *schema.OccupationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOccupationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}