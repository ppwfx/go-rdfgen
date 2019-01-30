package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDataFeedItem strings.Replacer
var strconvDataFeedItem strconv.NumError

var basicDataFeedItemTraitMapping = map[string]func(*schema.DataFeedItemTrait, []string){}
var customDataFeedItemTraitMapping = map[string]func(*schema.DataFeedItemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DataFeedItem", func(ctx jsonld.Context) (interface{}) {
		return NewDataFeedItemFromContext(ctx)
	})






	customDataFeedItemTraitMapping["http://schema.org/dateCreated"] = func(x *schema.DataFeedItemTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DateCreated, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DateCreated = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DateCreated = s
		}
	}

	customDataFeedItemTraitMapping["http://schema.org/dateModified"] = func(x *schema.DataFeedItemTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DateModified, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DateModified = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DateModified = s
		}
	}

	customDataFeedItemTraitMapping["http://schema.org/dateDeleted"] = func(x *schema.DataFeedItemTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DateDeleted = &y

		return
	}

	customDataFeedItemTraitMapping["http://schema.org/item"] = func(x *schema.DataFeedItemTrait, ctx jsonld.Context, s string){
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
}

func NewDataFeedItemFromContext(ctx jsonld.Context) (x schema.DataFeedItem) {
	x.Type = "http://schema.org/DataFeedItem"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDataFeedItemTrait(ctx, &x.DataFeedItemTrait)
	MapCustomToDataFeedItemTrait(ctx, &x.DataFeedItemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDataFeedItemTrait(ctx jsonld.Context, x *schema.DataFeedItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDataFeedItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDataFeedItemTrait(ctx jsonld.Context, x *schema.DataFeedItemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDataFeedItemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}