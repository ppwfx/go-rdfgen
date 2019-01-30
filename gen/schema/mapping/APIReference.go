package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAPIReference strings.Replacer
var strconvAPIReference strconv.NumError

var basicAPIReferenceTraitMapping = map[string]func(*schema.APIReferenceTrait, []string){}
var customAPIReferenceTraitMapping = map[string]func(*schema.APIReferenceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/APIReference", func(ctx jsonld.Context) (interface{}) {
		return NewAPIReferenceFromContext(ctx)
	})


	basicAPIReferenceTraitMapping["http://schema.org/assembly"] = func(x *schema.APIReferenceTrait, s []string) {
		x.Assembly = s[0]
	}


	basicAPIReferenceTraitMapping["http://schema.org/assemblyVersion"] = func(x *schema.APIReferenceTrait, s []string) {
		x.AssemblyVersion = s[0]
	}


	basicAPIReferenceTraitMapping["http://schema.org/programmingModel"] = func(x *schema.APIReferenceTrait, s []string) {
		x.ProgrammingModel = s[0]
	}


	basicAPIReferenceTraitMapping["http://schema.org/targetPlatform"] = func(x *schema.APIReferenceTrait, s []string) {
		x.TargetPlatform = s[0]
	}


	basicAPIReferenceTraitMapping["http://schema.org/executableLibraryName"] = func(x *schema.APIReferenceTrait, s []string) {
		x.ExecutableLibraryName = s[0]
	}

}

func NewAPIReferenceFromContext(ctx jsonld.Context) (x schema.APIReference) {
	x.Type = "http://schema.org/APIReference"
	MapBasicToTechArticleTrait(ctx, &x.TechArticleTrait)
	MapCustomToTechArticleTrait(ctx, &x.TechArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAPIReferenceTrait(ctx, &x.APIReferenceTrait)
	MapCustomToAPIReferenceTrait(ctx, &x.APIReferenceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAPIReferenceTrait(ctx jsonld.Context, x *schema.APIReferenceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAPIReferenceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAPIReferenceTrait(ctx jsonld.Context, x *schema.APIReferenceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAPIReferenceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}