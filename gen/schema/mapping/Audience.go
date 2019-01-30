package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAudience strings.Replacer
var strconvAudience strconv.NumError

var basicAudienceTraitMapping = map[string]func(*schema.AudienceTrait, []string){}
var customAudienceTraitMapping = map[string]func(*schema.AudienceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Audience", func(ctx jsonld.Context) (interface{}) {
		return NewAudienceFromContext(ctx)
	})


	basicAudienceTraitMapping["http://schema.org/audienceType"] = func(x *schema.AudienceTrait, s []string) {
		x.AudienceType = s[0]
	}



	customAudienceTraitMapping["http://schema.org/geographicArea"] = func(x *schema.AudienceTrait, ctx jsonld.Context, s string){
		var y schema.AdministrativeArea
		if strings.HasPrefix(s, "_:") {
			y = NewAdministrativeAreaFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAdministrativeArea()
			y.Id = s
		}

		x.GeographicArea = &y

		return
	}
}

func NewAudienceFromContext(ctx jsonld.Context) (x schema.Audience) {
	x.Type = "http://schema.org/Audience"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAudienceTrait(ctx, &x.AudienceTrait)
	MapCustomToAudienceTrait(ctx, &x.AudienceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAudienceTrait(ctx jsonld.Context, x *schema.AudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAudienceTrait(ctx jsonld.Context, x *schema.AudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}