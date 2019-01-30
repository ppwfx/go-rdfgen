package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOwnershipInfo strings.Replacer
var strconvOwnershipInfo strconv.NumError

var basicOwnershipInfoTraitMapping = map[string]func(*schema.OwnershipInfoTrait, []string){}
var customOwnershipInfoTraitMapping = map[string]func(*schema.OwnershipInfoTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OwnershipInfo", func(ctx jsonld.Context) (interface{}) {
		return NewOwnershipInfoFromContext(ctx)
	})






	customOwnershipInfoTraitMapping["http://schema.org/typeOfGood"] = func(x *schema.OwnershipInfoTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TypeOfGood, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TypeOfGood = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TypeOfGood = s
		}
	}

	customOwnershipInfoTraitMapping["http://schema.org/acquiredFrom"] = func(x *schema.OwnershipInfoTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AcquiredFrom, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AcquiredFrom = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AcquiredFrom = s
		}
	}

	customOwnershipInfoTraitMapping["http://schema.org/ownedFrom"] = func(x *schema.OwnershipInfoTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.OwnedFrom = &y

		return
	}

	customOwnershipInfoTraitMapping["http://schema.org/ownedThrough"] = func(x *schema.OwnershipInfoTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.OwnedThrough = &y

		return
	}
}

func NewOwnershipInfoFromContext(ctx jsonld.Context) (x schema.OwnershipInfo) {
	x.Type = "http://schema.org/OwnershipInfo"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOwnershipInfoTrait(ctx, &x.OwnershipInfoTrait)
	MapCustomToOwnershipInfoTrait(ctx, &x.OwnershipInfoTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOwnershipInfoTrait(ctx jsonld.Context, x *schema.OwnershipInfoTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOwnershipInfoTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOwnershipInfoTrait(ctx jsonld.Context, x *schema.OwnershipInfoTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOwnershipInfoTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}