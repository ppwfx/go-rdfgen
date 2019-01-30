package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLocksmith strings.Replacer
var strconvLocksmith strconv.NumError

var basicLocksmithTraitMapping = map[string]func(*schema.LocksmithTrait, []string){}
var customLocksmithTraitMapping = map[string]func(*schema.LocksmithTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Locksmith", func(ctx jsonld.Context) (interface{}) {
		return NewLocksmithFromContext(ctx)
	})

}

func NewLocksmithFromContext(ctx jsonld.Context) (x schema.Locksmith) {
	x.Type = "http://schema.org/Locksmith"
	MapBasicToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)
	MapCustomToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToLocksmithTrait(ctx, &x.LocksmithTrait)
	MapCustomToLocksmithTrait(ctx, &x.LocksmithTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLocksmithTrait(ctx jsonld.Context, x *schema.LocksmithTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLocksmithTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLocksmithTrait(ctx jsonld.Context, x *schema.LocksmithTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLocksmithTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}