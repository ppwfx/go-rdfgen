package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAnswer strings.Replacer
var strconvAnswer strconv.NumError

var basicAnswerTraitMapping = map[string]func(*schema.AnswerTrait, []string){}
var customAnswerTraitMapping = map[string]func(*schema.AnswerTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Answer", func(ctx jsonld.Context) (interface{}) {
		return NewAnswerFromContext(ctx)
	})

}

func NewAnswerFromContext(ctx jsonld.Context) (x schema.Answer) {
	x.Type = "http://schema.org/Answer"
	MapBasicToCommentTrait(ctx, &x.CommentTrait)
	MapCustomToCommentTrait(ctx, &x.CommentTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAnswerTrait(ctx, &x.AnswerTrait)
	MapCustomToAnswerTrait(ctx, &x.AnswerTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAnswerTrait(ctx jsonld.Context, x *schema.AnswerTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAnswerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAnswerTrait(ctx jsonld.Context, x *schema.AnswerTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAnswerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}