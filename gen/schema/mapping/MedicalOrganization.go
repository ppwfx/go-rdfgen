package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalOrganization strings.Replacer
var strconvMedicalOrganization strconv.NumError

var basicMedicalOrganizationTraitMapping = map[string]func(*schema.MedicalOrganizationTrait, []string){}
var customMedicalOrganizationTraitMapping = map[string]func(*schema.MedicalOrganizationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalOrganization", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalOrganizationFromContext(ctx)
	})


	basicMedicalOrganizationTraitMapping["http://schema.org/healthPlanNetworkId"] = func(x *schema.MedicalOrganizationTrait, s []string) {
		x.HealthPlanNetworkId = s[0]
	}


	basicMedicalOrganizationTraitMapping["http://schema.org/isAcceptingNewPatients"] = func(x *schema.MedicalOrganizationTrait, s []string) {
		var err error
		x.IsAcceptingNewPatients, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}



	customMedicalOrganizationTraitMapping["http://schema.org/medicalSpecialty"] = func(x *schema.MedicalOrganizationTrait, ctx jsonld.Context, s string){
		var y schema.MedicalSpecialty
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalSpecialtyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalSpecialty()
			y.Id = s
		}

		x.MedicalSpecialty = &y

		return
	}
}

func NewMedicalOrganizationFromContext(ctx jsonld.Context) (x schema.MedicalOrganization) {
	x.Type = "http://schema.org/MedicalOrganization"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)
	MapCustomToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalOrganizationTrait(ctx jsonld.Context, x *schema.MedicalOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalOrganizationTrait(ctx jsonld.Context, x *schema.MedicalOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}