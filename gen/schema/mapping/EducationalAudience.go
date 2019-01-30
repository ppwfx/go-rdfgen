package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEducationalAudience strings.Replacer
var strconvEducationalAudience strconv.NumError

var basicEducationalAudienceTraitMapping = map[string]func(*schema.EducationalAudienceTrait, []string){}
var customEducationalAudienceTraitMapping = map[string]func(*schema.EducationalAudienceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EducationalAudience", func(ctx jsonld.Context) (interface{}) {
		return NewEducationalAudienceFromContext(ctx)
	})


	basicEducationalAudienceTraitMapping["http://schema.org/educationalRole"] = func(x *schema.EducationalAudienceTrait, s []string) {
		x.EducationalRole = s[0]
	}

}

func NewEducationalAudienceFromContext(ctx jsonld.Context) (x schema.EducationalAudience) {
	x.Type = "http://schema.org/EducationalAudience"
	MapBasicToAudienceTrait(ctx, &x.AudienceTrait)
	MapCustomToAudienceTrait(ctx, &x.AudienceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEducationalAudienceTrait(ctx, &x.EducationalAudienceTrait)
	MapCustomToEducationalAudienceTrait(ctx, &x.EducationalAudienceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEducationalAudienceTrait(ctx jsonld.Context, x *schema.EducationalAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEducationalAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEducationalAudienceTrait(ctx jsonld.Context, x *schema.EducationalAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEducationalAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}