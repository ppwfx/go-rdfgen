package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEducationalOrganization strings.Replacer
var strconvEducationalOrganization strconv.NumError

var basicEducationalOrganizationTraitMapping = map[string]func(*schema.EducationalOrganizationTrait, []string){}
var customEducationalOrganizationTraitMapping = map[string]func(*schema.EducationalOrganizationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EducationalOrganization", func(ctx jsonld.Context) (interface{}) {
		return NewEducationalOrganizationFromContext(ctx)
	})



	customEducationalOrganizationTraitMapping["http://schema.org/alumni"] = func(x *schema.EducationalOrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Alumni = &y

		return
	}
}

func NewEducationalOrganizationFromContext(ctx jsonld.Context) (x schema.EducationalOrganization) {
	x.Type = "http://schema.org/EducationalOrganization"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)
	MapCustomToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEducationalOrganizationTrait(ctx jsonld.Context, x *schema.EducationalOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEducationalOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEducationalOrganizationTrait(ctx jsonld.Context, x *schema.EducationalOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEducationalOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}