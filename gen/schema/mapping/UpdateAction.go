package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUpdateAction strings.Replacer
var strconvUpdateAction strconv.NumError

var basicUpdateActionTraitMapping = map[string]func(*schema.UpdateActionTrait, []string){}
var customUpdateActionTraitMapping = map[string]func(*schema.UpdateActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UpdateAction", func(ctx jsonld.Context) (interface{}) {
		return NewUpdateActionFromContext(ctx)
	})




	customUpdateActionTraitMapping["http://schema.org/targetCollection"] = func(x *schema.UpdateActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.TargetCollection = &y

		return
	}

	customUpdateActionTraitMapping["http://schema.org/collection"] = func(x *schema.UpdateActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Collection = &y

		return
	}
}

func NewUpdateActionFromContext(ctx jsonld.Context) (x schema.UpdateAction) {
	x.Type = "http://schema.org/UpdateAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUpdateActionTrait(ctx, &x.UpdateActionTrait)
	MapCustomToUpdateActionTrait(ctx, &x.UpdateActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUpdateActionTrait(ctx jsonld.Context, x *schema.UpdateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUpdateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUpdateActionTrait(ctx jsonld.Context, x *schema.UpdateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUpdateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}