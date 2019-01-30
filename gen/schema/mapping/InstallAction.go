package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInstallAction strings.Replacer
var strconvInstallAction strconv.NumError

var basicInstallActionTraitMapping = map[string]func(*schema.InstallActionTrait, []string){}
var customInstallActionTraitMapping = map[string]func(*schema.InstallActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InstallAction", func(ctx jsonld.Context) (interface{}) {
		return NewInstallActionFromContext(ctx)
	})

}

func NewInstallActionFromContext(ctx jsonld.Context) (x schema.InstallAction) {
	x.Type = "http://schema.org/InstallAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInstallActionTrait(ctx, &x.InstallActionTrait)
	MapCustomToInstallActionTrait(ctx, &x.InstallActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInstallActionTrait(ctx jsonld.Context, x *schema.InstallActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInstallActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInstallActionTrait(ctx jsonld.Context, x *schema.InstallActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInstallActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}