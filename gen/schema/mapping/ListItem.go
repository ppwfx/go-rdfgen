package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsListItem strings.Replacer
var strconvListItem strconv.NumError

var basicListItemTraitMapping = map[string]func(*schema.ListItemTrait, []string){}
var customListItemTraitMapping = map[string]func(*schema.ListItemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ListItem", func(ctx jsonld.Context) (interface{}) {
		return NewListItemFromContext(ctx)
	})






	customListItemTraitMapping["http://schema.org/position"] = func(x *schema.ListItemTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Position, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Position = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Position = s
		}
	}

	customListItemTraitMapping["http://schema.org/item"] = func(x *schema.ListItemTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Item = &y

		return
	}

	customListItemTraitMapping["http://schema.org/nextItem"] = func(x *schema.ListItemTrait, ctx jsonld.Context, s string){
		var y schema.ListItem
		if strings.HasPrefix(s, "_:") {
			y = NewListItemFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewListItem()
			y.Id = s
		}

		x.NextItem = &y

		return
	}

	customListItemTraitMapping["http://schema.org/previousItem"] = func(x *schema.ListItemTrait, ctx jsonld.Context, s string){
		var y schema.ListItem
		if strings.HasPrefix(s, "_:") {
			y = NewListItemFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewListItem()
			y.Id = s
		}

		x.PreviousItem = &y

		return
	}
}

func NewListItemFromContext(ctx jsonld.Context) (x schema.ListItem) {
	x.Type = "http://schema.org/ListItem"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToListItemTrait(ctx jsonld.Context, x *schema.ListItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicListItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToListItemTrait(ctx jsonld.Context, x *schema.ListItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customListItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}