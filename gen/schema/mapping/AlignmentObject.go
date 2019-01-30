package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAlignmentObject strings.Replacer
var strconvAlignmentObject strconv.NumError

var basicAlignmentObjectTraitMapping = map[string]func(*schema.AlignmentObjectTrait, []string){}
var customAlignmentObjectTraitMapping = map[string]func(*schema.AlignmentObjectTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AlignmentObject", func(ctx jsonld.Context) (interface{}) {
		return NewAlignmentObjectFromContext(ctx)
	})


	basicAlignmentObjectTraitMapping["http://schema.org/alignmentType"] = func(x *schema.AlignmentObjectTrait, s []string) {
		x.AlignmentType = s[0]
	}


	basicAlignmentObjectTraitMapping["http://schema.org/educationalFramework"] = func(x *schema.AlignmentObjectTrait, s []string) {
		x.EducationalFramework = s[0]
	}


	basicAlignmentObjectTraitMapping["http://schema.org/targetDescription"] = func(x *schema.AlignmentObjectTrait, s []string) {
		x.TargetDescription = s[0]
	}


	basicAlignmentObjectTraitMapping["http://schema.org/targetName"] = func(x *schema.AlignmentObjectTrait, s []string) {
		x.TargetName = s[0]
	}


	basicAlignmentObjectTraitMapping["http://schema.org/targetUrl"] = func(x *schema.AlignmentObjectTrait, s []string) {
		x.TargetUrl = s[0]
	}

}

func NewAlignmentObjectFromContext(ctx jsonld.Context) (x schema.AlignmentObject) {
	x.Type = "http://schema.org/AlignmentObject"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAlignmentObjectTrait(ctx, &x.AlignmentObjectTrait)
	MapCustomToAlignmentObjectTrait(ctx, &x.AlignmentObjectTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAlignmentObjectTrait(ctx jsonld.Context, x *schema.AlignmentObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAlignmentObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAlignmentObjectTrait(ctx jsonld.Context, x *schema.AlignmentObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAlignmentObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}