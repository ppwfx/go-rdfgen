package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReplaceAction strings.Replacer
var strconvReplaceAction strconv.NumError

var basicReplaceActionTraitMapping = map[string]func(*schema.ReplaceActionTrait, []string){}
var customReplaceActionTraitMapping = map[string]func(*schema.ReplaceActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReplaceAction", func(ctx jsonld.Context) (interface{}) {
		return NewReplaceActionFromContext(ctx)
	})




	customReplaceActionTraitMapping["http://schema.org/replacer"] = func(x *schema.ReplaceActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Replacer = &y

		return
	}

	customReplaceActionTraitMapping["http://schema.org/replacee"] = func(x *schema.ReplaceActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Replacee = &y

		return
	}
}

func NewReplaceActionFromContext(ctx jsonld.Context) (x schema.ReplaceAction) {
	x.Type = "http://schema.org/ReplaceAction"
	MapBasicToUpdateActionTrait(ctx, &x.UpdateActionTrait)
	MapCustomToUpdateActionTrait(ctx, &x.UpdateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReplaceActionTrait(ctx, &x.ReplaceActionTrait)
	MapCustomToReplaceActionTrait(ctx, &x.ReplaceActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReplaceActionTrait(ctx jsonld.Context, x *schema.ReplaceActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReplaceActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReplaceActionTrait(ctx jsonld.Context, x *schema.ReplaceActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReplaceActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}