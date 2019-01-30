package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQuotation strings.Replacer
var strconvQuotation strconv.NumError

var basicQuotationTraitMapping = map[string]func(*schema.QuotationTrait, []string){}
var customQuotationTraitMapping = map[string]func(*schema.QuotationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Quotation", func(ctx jsonld.Context) (interface{}) {
		return NewQuotationFromContext(ctx)
	})



	customQuotationTraitMapping["http://schema.org/spokenByCharacter"] = func(x *schema.QuotationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SpokenByCharacter, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SpokenByCharacter = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SpokenByCharacter = s
		}
	}
}

func NewQuotationFromContext(ctx jsonld.Context) (x schema.Quotation) {
	x.Type = "http://schema.org/Quotation"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQuotationTrait(ctx, &x.QuotationTrait)
	MapCustomToQuotationTrait(ctx, &x.QuotationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQuotationTrait(ctx jsonld.Context, x *schema.QuotationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQuotationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQuotationTrait(ctx jsonld.Context, x *schema.QuotationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQuotationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}