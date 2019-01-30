package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalClinic strings.Replacer
var strconvMedicalClinic strconv.NumError

var basicMedicalClinicTraitMapping = map[string]func(*schema.MedicalClinicTrait, []string){}
var customMedicalClinicTraitMapping = map[string]func(*schema.MedicalClinicTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalClinic", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalClinicFromContext(ctx)
	})




	customMedicalClinicTraitMapping["http://schema.org/medicalSpecialty"] = func(x *schema.MedicalClinicTrait, ctx jsonld.Context, s string){
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

	customMedicalClinicTraitMapping["http://schema.org/availableService"] = func(x *schema.MedicalClinicTrait, ctx jsonld.Context, s string){
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

func NewMedicalClinicFromContext(ctx jsonld.Context) (x schema.MedicalClinic) {
	x.Type = "http://schema.org/MedicalClinic"
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


	MapBasicToMedicalClinicTrait(ctx, &x.MedicalClinicTrait)
	MapCustomToMedicalClinicTrait(ctx, &x.MedicalClinicTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalClinicTrait(ctx jsonld.Context, x *schema.MedicalClinicTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalClinicTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalClinicTrait(ctx jsonld.Context, x *schema.MedicalClinicTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalClinicTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}