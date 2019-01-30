package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPhysician strings.Replacer
var strconvPhysician strconv.NumError

var basicPhysicianTraitMapping = map[string]func(*schema.PhysicianTrait, []string){}
var customPhysicianTraitMapping = map[string]func(*schema.PhysicianTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Physician", func(ctx jsonld.Context) (interface{}) {
		return NewPhysicianFromContext(ctx)
	})





	customPhysicianTraitMapping["http://schema.org/medicalSpecialty"] = func(x *schema.PhysicianTrait, ctx jsonld.Context, s string){
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

	customPhysicianTraitMapping["http://schema.org/availableService"] = func(x *schema.PhysicianTrait, ctx jsonld.Context, s string){
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

	customPhysicianTraitMapping["http://schema.org/hospitalAffiliation"] = func(x *schema.PhysicianTrait, ctx jsonld.Context, s string){
		var y schema.Hospital
		if strings.HasPrefix(s, "_:") {
			y = NewHospitalFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewHospital()
			y.Id = s
		}

		x.HospitalAffiliation = &y

		return
	}
}

func NewPhysicianFromContext(ctx jsonld.Context) (x schema.Physician) {
	x.Type = "http://schema.org/Physician"
	MapBasicToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)
	MapCustomToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)
	MapCustomToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)


	MapBasicToPhysicianTrait(ctx, &x.PhysicianTrait)
	MapCustomToPhysicianTrait(ctx, &x.PhysicianTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPhysicianTrait(ctx jsonld.Context, x *schema.PhysicianTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPhysicianTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPhysicianTrait(ctx jsonld.Context, x *schema.PhysicianTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPhysicianTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}