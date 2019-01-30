package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsParentAudience strings.Replacer
var strconvParentAudience strconv.NumError

var basicParentAudienceTraitMapping = map[string]func(*schema.ParentAudienceTrait, []string){}
var customParentAudienceTraitMapping = map[string]func(*schema.ParentAudienceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ParentAudience", func(ctx jsonld.Context) (interface{}) {
		return NewParentAudienceFromContext(ctx)
	})


	basicParentAudienceTraitMapping["http://schema.org/childMaxAge"] = func(x *schema.ParentAudienceTrait, s []string) {
		var err error
		x.ChildMaxAge, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicParentAudienceTraitMapping["http://schema.org/childMinAge"] = func(x *schema.ParentAudienceTrait, s []string) {
		var err error
		x.ChildMinAge, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}

}

func NewParentAudienceFromContext(ctx jsonld.Context) (x schema.ParentAudience) {
	x.Type = "http://schema.org/ParentAudience"
	MapBasicToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)
	MapCustomToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)

	MapBasicToAudienceTrait(ctx, &x.AudienceTrait)
	MapCustomToAudienceTrait(ctx, &x.AudienceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToParentAudienceTrait(ctx, &x.ParentAudienceTrait)
	MapCustomToParentAudienceTrait(ctx, &x.ParentAudienceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToParentAudienceTrait(ctx jsonld.Context, x *schema.ParentAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicParentAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToParentAudienceTrait(ctx jsonld.Context, x *schema.ParentAudienceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customParentAudienceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}