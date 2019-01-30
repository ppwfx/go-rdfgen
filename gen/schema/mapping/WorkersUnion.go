package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWorkersUnion strings.Replacer
var strconvWorkersUnion strconv.NumError

var basicWorkersUnionTraitMapping = map[string]func(*schema.WorkersUnionTrait, []string){}
var customWorkersUnionTraitMapping = map[string]func(*schema.WorkersUnionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WorkersUnion", func(ctx jsonld.Context) (interface{}) {
		return NewWorkersUnionFromContext(ctx)
	})

}

func NewWorkersUnionFromContext(ctx jsonld.Context) (x schema.WorkersUnion) {
	x.Type = "http://schema.org/WorkersUnion"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWorkersUnionTrait(ctx, &x.WorkersUnionTrait)
	MapCustomToWorkersUnionTrait(ctx, &x.WorkersUnionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWorkersUnionTrait(ctx jsonld.Context, x *schema.WorkersUnionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWorkersUnionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWorkersUnionTrait(ctx jsonld.Context, x *schema.WorkersUnionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWorkersUnionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}