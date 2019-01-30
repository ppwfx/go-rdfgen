package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHospital strings.Replacer
var strconvHospital strconv.NumError

var basicHospitalTraitMapping = map[string]func(*schema.HospitalTrait, []string){}
var customHospitalTraitMapping = map[string]func(*schema.HospitalTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Hospital", func(ctx jsonld.Context) (interface{}) {
		return NewHospitalFromContext(ctx)
	})




	customHospitalTraitMapping["http://schema.org/medicalSpecialty"] = func(x *schema.HospitalTrait, ctx jsonld.Context, s string){
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

	customHospitalTraitMapping["http://schema.org/availableService"] = func(x *schema.HospitalTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AvailableService, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AvailableService = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AvailableService = s
		}
	}
}

func NewHospitalFromContext(ctx jsonld.Context) (x schema.Hospital) {
	x.Type = "http://schema.org/Hospital"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)
	MapCustomToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToEmergencyServiceTrait(ctx, &x.EmergencyServiceTrait)
	MapCustomToEmergencyServiceTrait(ctx, &x.EmergencyServiceTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)


	MapBasicToHospitalTrait(ctx, &x.HospitalTrait)
	MapCustomToHospitalTrait(ctx, &x.HospitalTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHospitalTrait(ctx jsonld.Context, x *schema.HospitalTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHospitalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHospitalTrait(ctx jsonld.Context, x *schema.HospitalTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHospitalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}