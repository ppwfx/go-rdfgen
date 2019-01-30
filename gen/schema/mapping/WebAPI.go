package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWebAPI strings.Replacer
var strconvWebAPI strconv.NumError

var basicWebAPITraitMapping = map[string]func(*schema.WebAPITrait, []string){}
var customWebAPITraitMapping = map[string]func(*schema.WebAPITrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WebAPI", func(ctx jsonld.Context) (interface{}) {
		return NewWebAPIFromContext(ctx)
	})



	customWebAPITraitMapping["http://schema.org/documentation"] = func(x *schema.WebAPITrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Documentation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Documentation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Documentation = s
		}
	}
}

func NewWebAPIFromContext(ctx jsonld.Context) (x schema.WebAPI) {
	x.Type = "http://schema.org/WebAPI"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWebAPITrait(ctx, &x.WebAPITrait)
	MapCustomToWebAPITrait(ctx, &x.WebAPITrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWebAPITrait(ctx jsonld.Context, x *schema.WebAPITrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWebAPITraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWebAPITrait(ctx jsonld.Context, x *schema.WebAPITrait) () {
	for k, v := range ctx.Current {
		f, ok := customWebAPITraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}