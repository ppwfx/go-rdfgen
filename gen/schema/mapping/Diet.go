package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDiet strings.Replacer
var strconvDiet strconv.NumError

var basicDietTraitMapping = map[string]func(*schema.DietTrait, []string){}
var customDietTraitMapping = map[string]func(*schema.DietTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Diet", func(ctx jsonld.Context) (interface{}) {
		return NewDietFromContext(ctx)
	})


	basicDietTraitMapping["http://schema.org/expertConsiderations"] = func(x *schema.DietTrait, s []string) {
		x.ExpertConsiderations = s[0]
	}


	basicDietTraitMapping["http://schema.org/overview"] = func(x *schema.DietTrait, s []string) {
		x.Overview = s[0]
	}


	basicDietTraitMapping["http://schema.org/dietFeatures"] = func(x *schema.DietTrait, s []string) {
		x.DietFeatures = s[0]
	}


	basicDietTraitMapping["http://schema.org/physiologicalBenefits"] = func(x *schema.DietTrait, s []string) {
		x.PhysiologicalBenefits = s[0]
	}


	basicDietTraitMapping["http://schema.org/risks"] = func(x *schema.DietTrait, s []string) {
		x.Risks = s[0]
	}



	customDietTraitMapping["http://schema.org/endorsers"] = func(x *schema.DietTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Endorsers, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Endorsers = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Endorsers = s
		}
	}
}

func NewDietFromContext(ctx jsonld.Context) (x schema.Diet) {
	x.Type = "http://schema.org/Diet"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)
	MapCustomToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)


	MapBasicToDietTrait(ctx, &x.DietTrait)
	MapCustomToDietTrait(ctx, &x.DietTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDietTrait(ctx jsonld.Context, x *schema.DietTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDietTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDietTrait(ctx jsonld.Context, x *schema.DietTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDietTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}