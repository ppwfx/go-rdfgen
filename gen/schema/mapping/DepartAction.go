package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDepartAction strings.Replacer
var strconvDepartAction strconv.NumError

var basicDepartActionTraitMapping = map[string]func(*schema.DepartActionTrait, []string){}
var customDepartActionTraitMapping = map[string]func(*schema.DepartActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DepartAction", func(ctx jsonld.Context) (interface{}) {
		return NewDepartActionFromContext(ctx)
	})

}

func NewDepartActionFromContext(ctx jsonld.Context) (x schema.DepartAction) {
	x.Type = "http://schema.org/DepartAction"
	MapBasicToMoveActionTrait(ctx, &x.MoveActionTrait)
	MapCustomToMoveActionTrait(ctx, &x.MoveActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDepartActionTrait(ctx, &x.DepartActionTrait)
	MapCustomToDepartActionTrait(ctx, &x.DepartActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDepartActionTrait(ctx jsonld.Context, x *schema.DepartActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDepartActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDepartActionTrait(ctx jsonld.Context, x *schema.DepartActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDepartActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}