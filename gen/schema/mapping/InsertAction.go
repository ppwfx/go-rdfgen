package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInsertAction strings.Replacer
var strconvInsertAction strconv.NumError

var basicInsertActionTraitMapping = map[string]func(*schema.InsertActionTrait, []string){}
var customInsertActionTraitMapping = map[string]func(*schema.InsertActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InsertAction", func(ctx jsonld.Context) (interface{}) {
		return NewInsertActionFromContext(ctx)
	})



	customInsertActionTraitMapping["http://schema.org/toLocation"] = func(x *schema.InsertActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ToLocation = &y

		return
	}
}

func NewInsertActionFromContext(ctx jsonld.Context) (x schema.InsertAction) {
	x.Type = "http://schema.org/InsertAction"
	MapBasicToAddActionTrait(ctx, &x.AddActionTrait)
	MapCustomToAddActionTrait(ctx, &x.AddActionTrait)

	MapBasicToUpdateActionTrait(ctx, &x.UpdateActionTrait)
	MapCustomToUpdateActionTrait(ctx, &x.UpdateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInsertActionTrait(ctx, &x.InsertActionTrait)
	MapCustomToInsertActionTrait(ctx, &x.InsertActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInsertActionTrait(ctx jsonld.Context, x *schema.InsertActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInsertActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInsertActionTrait(ctx jsonld.Context, x *schema.InsertActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInsertActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}