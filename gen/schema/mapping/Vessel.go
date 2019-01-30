package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVessel strings.Replacer
var strconvVessel strconv.NumError

var basicVesselTraitMapping = map[string]func(*schema.VesselTrait, []string){}
var customVesselTraitMapping = map[string]func(*schema.VesselTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Vessel", func(ctx jsonld.Context) (interface{}) {
		return NewVesselFromContext(ctx)
	})

}

func NewVesselFromContext(ctx jsonld.Context) (x schema.Vessel) {
	x.Type = "http://schema.org/Vessel"
	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVesselTrait(ctx, &x.VesselTrait)
	MapCustomToVesselTrait(ctx, &x.VesselTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVesselTrait(ctx jsonld.Context, x *schema.VesselTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVesselTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVesselTrait(ctx jsonld.Context, x *schema.VesselTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVesselTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}