package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAppendAction strings.Replacer
var strconvAppendAction strconv.NumError

var basicAppendActionTraitMapping = map[string]func(*schema.AppendActionTrait, []string){}
var customAppendActionTraitMapping = map[string]func(*schema.AppendActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AppendAction", func(ctx jsonld.Context) (interface{}) {
		return NewAppendActionFromContext(ctx)
	})

}

func NewAppendActionFromContext(ctx jsonld.Context) (x schema.AppendAction) {
	x.Type = "http://schema.org/AppendAction"
	MapBasicToInsertActionTrait(ctx, &x.InsertActionTrait)
	MapCustomToInsertActionTrait(ctx, &x.InsertActionTrait)

	MapBasicToAddActionTrait(ctx, &x.AddActionTrait)
	MapCustomToAddActionTrait(ctx, &x.AddActionTrait)

	MapBasicToUpdateActionTrait(ctx, &x.UpdateActionTrait)
	MapCustomToUpdateActionTrait(ctx, &x.UpdateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAppendActionTrait(ctx, &x.AppendActionTrait)
	MapCustomToAppendActionTrait(ctx, &x.AppendActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAppendActionTrait(ctx jsonld.Context, x *schema.AppendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAppendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAppendActionTrait(ctx jsonld.Context, x *schema.AppendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAppendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}