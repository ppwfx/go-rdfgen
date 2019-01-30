package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsConversation strings.Replacer
var strconvConversation strconv.NumError

var basicConversationTraitMapping = map[string]func(*schema.ConversationTrait, []string){}
var customConversationTraitMapping = map[string]func(*schema.ConversationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Conversation", func(ctx jsonld.Context) (interface{}) {
		return NewConversationFromContext(ctx)
	})

}

func NewConversationFromContext(ctx jsonld.Context) (x schema.Conversation) {
	x.Type = "http://schema.org/Conversation"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToConversationTrait(ctx, &x.ConversationTrait)
	MapCustomToConversationTrait(ctx, &x.ConversationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToConversationTrait(ctx jsonld.Context, x *schema.ConversationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicConversationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToConversationTrait(ctx jsonld.Context, x *schema.ConversationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customConversationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}