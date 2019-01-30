package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCategoryCode strings.Replacer
var strconvCategoryCode strconv.NumError

var basicCategoryCodeTraitMapping = map[string]func(*schema.CategoryCodeTrait, []string){}
var customCategoryCodeTraitMapping = map[string]func(*schema.CategoryCodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CategoryCode", func(ctx jsonld.Context) (interface{}) {
		return NewCategoryCodeFromContext(ctx)
	})


	basicCategoryCodeTraitMapping["http://schema.org/codeValue"] = func(x *schema.CategoryCodeTrait, s []string) {
		x.CodeValue = s[0]
	}



	customCategoryCodeTraitMapping["http://schema.org/inCodeSet"] = func(x *schema.CategoryCodeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.InCodeSet, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.InCodeSet = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.InCodeSet = s
		}
	}
}

func NewCategoryCodeFromContext(ctx jsonld.Context) (x schema.CategoryCode) {
	x.Type = "http://schema.org/CategoryCode"
	MapBasicToDefinedTermTrait(ctx, &x.DefinedTermTrait)
	MapCustomToDefinedTermTrait(ctx, &x.DefinedTermTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCategoryCodeTrait(ctx, &x.CategoryCodeTrait)
	MapCustomToCategoryCodeTrait(ctx, &x.CategoryCodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCategoryCodeTrait(ctx jsonld.Context, x *schema.CategoryCodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCategoryCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCategoryCodeTrait(ctx jsonld.Context, x *schema.CategoryCodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCategoryCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}