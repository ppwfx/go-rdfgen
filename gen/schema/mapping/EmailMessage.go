package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEmailMessage strings.Replacer
var strconvEmailMessage strconv.NumError

var basicEmailMessageTraitMapping = map[string]func(*schema.EmailMessageTrait, []string){}
var customEmailMessageTraitMapping = map[string]func(*schema.EmailMessageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EmailMessage", func(ctx jsonld.Context) (interface{}) {
		return NewEmailMessageFromContext(ctx)
	})

}

func NewEmailMessageFromContext(ctx jsonld.Context) (x schema.EmailMessage) {
	x.Type = "http://schema.org/EmailMessage"
	MapBasicToMessageTrait(ctx, &x.MessageTrait)
	MapCustomToMessageTrait(ctx, &x.MessageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEmailMessageTrait(ctx, &x.EmailMessageTrait)
	MapCustomToEmailMessageTrait(ctx, &x.EmailMessageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEmailMessageTrait(ctx jsonld.Context, x *schema.EmailMessageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEmailMessageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEmailMessageTrait(ctx jsonld.Context, x *schema.EmailMessageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEmailMessageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}