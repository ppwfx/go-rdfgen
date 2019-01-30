package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalStudy strings.Replacer
var strconvMedicalStudy strconv.NumError

var basicMedicalStudyTraitMapping = map[string]func(*schema.MedicalStudyTrait, []string){}
var customMedicalStudyTraitMapping = map[string]func(*schema.MedicalStudyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalStudy", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalStudyFromContext(ctx)
	})





	basicMedicalStudyTraitMapping["http://schema.org/population"] = func(x *schema.MedicalStudyTrait, s []string) {
		x.Population = s[0]
	}





	customMedicalStudyTraitMapping["http://schema.org/sponsor"] = func(x *schema.MedicalStudyTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Sponsor, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Sponsor = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Sponsor = s
		}
	}

	customMedicalStudyTraitMapping["http://schema.org/healthCondition"] = func(x *schema.MedicalStudyTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCondition
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCondition()
			y.Id = s
		}

		x.HealthCondition = &y

		return
	}

	customMedicalStudyTraitMapping["http://schema.org/outcome"] = func(x *schema.MedicalStudyTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Outcome, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Outcome = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Outcome = s
		}
	}

	customMedicalStudyTraitMapping["http://schema.org/status"] = func(x *schema.MedicalStudyTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Status, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Status = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Status = s
		}
	}

	customMedicalStudyTraitMapping["http://schema.org/studyLocation"] = func(x *schema.MedicalStudyTrait, ctx jsonld.Context, s string){
		var y schema.AdministrativeArea
		if strings.HasPrefix(s, "_:") {
			y = NewAdministrativeAreaFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAdministrativeArea()
			y.Id = s
		}

		x.StudyLocation = &y

		return
	}

	customMedicalStudyTraitMapping["http://schema.org/studySubject"] = func(x *schema.MedicalStudyTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.StudySubject = &y

		return
	}
}

func NewMedicalStudyFromContext(ctx jsonld.Context) (x schema.MedicalStudy) {
	x.Type = "http://schema.org/MedicalStudy"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalStudyTrait(ctx, &x.MedicalStudyTrait)
	MapCustomToMedicalStudyTrait(ctx, &x.MedicalStudyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalStudyTrait(ctx jsonld.Context, x *schema.MedicalStudyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalStudyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalStudyTrait(ctx jsonld.Context, x *schema.MedicalStudyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalStudyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}