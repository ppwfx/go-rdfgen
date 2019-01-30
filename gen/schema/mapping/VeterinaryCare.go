package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVeterinaryCare strings.Replacer
var strconvVeterinaryCare strconv.NumError

var basicVeterinaryCareTraitMapping = map[string]func(*schema.VeterinaryCareTrait, []string){}
var customVeterinaryCareTraitMapping = map[string]func(*schema.VeterinaryCareTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VeterinaryCare", func(ctx jsonld.Context) (interface{}) {
		return NewVeterinaryCareFromContext(ctx)
	})

}

func NewVeterinaryCareFromContext(ctx jsonld.Context) (x schema.VeterinaryCare) {
	x.Type = "http://schema.org/VeterinaryCare"
	MapBasicToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)
	MapCustomToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVeterinaryCareTrait(ctx, &x.VeterinaryCareTrait)
	MapCustomToVeterinaryCareTrait(ctx, &x.VeterinaryCareTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVeterinaryCareTrait(ctx jsonld.Context, x *schema.VeterinaryCareTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVeterinaryCareTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVeterinaryCareTrait(ctx jsonld.Context, x *schema.VeterinaryCareTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVeterinaryCareTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}